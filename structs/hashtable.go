package structs

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

const initSize = 3
const prime = 31
const loadThreshold = 0.75

type Hashtable struct {
	buckets []*LLNode
	length  int
	rnd     int
}

type HashCoder interface {
	HashCode() int
}

type item struct {
	key interface{}
	val interface{}
}

func NewHashtable() *Hashtable {
	randSrc := rand.NewSource(time.Now().Unix())
	randGen := rand.New(randSrc)

	return &Hashtable{
		buckets: make([]*LLNode, initSize),
		length:  0,
		rnd:     randGen.Int(),
	}
}

func (t *Hashtable) Length() int {
	if t == nil {
		return 0
	}
	return t.length
}

func (t *Hashtable) Get(key interface{}) interface{} {
	if t == nil || key == nil {
		return nil
	}

	index := t.bucketIndex(key)

	for node := t.buckets[index]; node != nil; node = node.Next {
		item := node.Data.(item)
		if item.key == key {
			return item.val
		}
	}
	return nil
}

func (t *Hashtable) Add(key interface{}, val interface{}) error {
	if t == nil {
		return fmt.Errorf("Failed access a nil hashtable")
	}
	if key == nil {
		return fmt.Errorf("hashtable key cannot be nil")
	}

	t.addInternal(key, val)

	load := float64(t.length) / float64(len(t.buckets))
	if load > loadThreshold {
		t.rehash()
	}

	return nil
}

func (t *Hashtable) Remove(key interface{}) error {
	if t == nil {
		return fmt.Errorf("Failed access a nil hashtable")
	}
	if key == nil {
		return fmt.Errorf("hashtable key cannot be nil")
	}

	i := t.bucketIndex(key)

	var prev *LLNode = nil

	node := t.buckets[i]
	for node != nil {
		item := node.Data.(item)
		if item.key == key {
			if prev == nil {
				t.buckets[i] = node.Next
			} else {
				prev.Next = node.Next
			}
			break
		}

		prev = node
		node = node.Next
	}

	return nil
}

func (t *Hashtable) addInternal(key interface{}, val interface{}) {
	i := t.bucketIndex(key)

	node := t.buckets[i]
	node = NewLLNode(node, item{key, val})
	t.buckets[i] = node

	t.length++
}

func (t *Hashtable) rehash() {
	buckets := t.buckets
	t.buckets = make([]*LLNode, len(buckets)*2)

	for _, node := range buckets {
		for ; node != nil; node = node.Next {
			item := node.Data.(item)
			t.addInternal(item.key, item.val)
		}
	}
}

func (t *Hashtable) bucketIndex(key interface{}) int {
	return t.hashCode(key) % len(t.buckets)
}

func (t *Hashtable) hashCode(key interface{}) int {
	var hash uint

	switch k := key.(type) {

	case HashCoder:
		hash = uint(k.HashCode())

	case string:
		hash = stringHash(k, t.rnd)

	case rune:
		hash = stringHash(string([]rune{k}), t.rnd)

	case uint:
		hash = numberHash(k, t.rnd)

	case int:
		hash = numberHash(uint(k), t.rnd)

	case uint8:
		hash = numberHash(uint(k), t.rnd)

	case uint16:
		hash = numberHash(uint(k), t.rnd)

	case uint32:
		hash = numberHash(uint(k), t.rnd)

	case uint64:
		hash = numberHash(uint(k), t.rnd)

	case uintptr:
		hash = numberHash(uint(k), t.rnd)
	}

	if unsafe.Sizeof(hash) == 4 {
		return int(hash & 0x7FFFFFFF)
	}

	return int(hash & 0x7FFFFFFFFFFFFFFF)
}

func stringHash(key string, rnd int) (hash uint) {
	hash = uint(rnd)
	for _, b := range []byte(key) {
		hash = prime*hash + uint(b)
	}
	return
}

func numberHash(key uint, rnd int) uint {
	return key * uint(rnd)
}
