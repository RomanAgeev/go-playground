package structs

import (
	"math/rand"
	"time"
	"unsafe"
)

const initialSize = 31
const prime = 31

type HashCoder interface {
	GetHashCode() int
}

type Hashtable struct {
	buckets []*LLNode
	length  int
	hash    hashCode
}

type hashCode32 struct {
	rnd uint32
}

type hashCode64 struct {
	rnd uint64
}

type hashCode interface {
	calHashCode(key interface{}) int
}

func (hash hashCode32) calHashCode(key interface{}) int {
	var h uint32
	switch k := key.(type) {
	case int:
		h = uint32(k) * hash.rnd
	case string:
		h = hash.rnd
		for _, b := range []byte(k) {
			h = prime*h + uint32(b)
		}
	case HashCoder:
		h = uint32(k.GetHashCode())
	default:
		panic("Cannot calc hashcode for the key")
	}

	return int(h & 0x7FFFFFFF)
}

func (hash hashCode64) calHashCode(key interface{}) int {
	var h uint64
	switch k := key.(type) {
	case int:
		h = uint64(k) * hash.rnd
	case string:
		h = hash.rnd
		for _, b := range []byte(k) {
			h = prime*h + uint64(b)
		}
	case HashCoder:
		h = uint64(k.GetHashCode())
	default:
		panic("Cannot calc hashcode for the key")
	}

	return int(h & 0x7FFFFFFFFFFFFFFF)
}

type hashItem struct {
	key interface{}
	val interface{}
}

func NewHashtable() *Hashtable {
	randSrc := rand.NewSource(time.Now().Unix())
	randGen := rand.New(randSrc)

	var hCode hashCode
	if unsafe.Sizeof(int(0)) == 8 {
		hCode = hashCode64{
			rnd: randGen.Uint64(),
		}
	} else {
		hCode = hashCode32{
			rnd: randGen.Uint32(),
		}
	}

	return &Hashtable{
		buckets: make([]*LLNode, initialSize),
		length:  0,
		hash:    hCode,
	}
}

func (table Hashtable) Length() int {
	return table.length
}

func (table Hashtable) Get(key interface{}) interface{} {
	if key == nil {
		panic("A hashtable key cannot be nil")
	}

	index := table.getBucketIndex(key)

	for node := table.buckets[index]; node != nil; node = node.Next {
		item := node.Data.(hashItem)
		if item.key == key {
			return item.val
		}
	}
	return nil
}

func (table *Hashtable) Add(key interface{}, val interface{}) {
	if table == nil {
		panic("A hashtable cannot be nil")
	}

	if key == nil {
		panic("A hashtable key cannot be nil")
	}

	index := table.getBucketIndex(key)

	node := table.buckets[index]
	node = NewLLNode(node, hashItem{key, val})
	table.buckets[index] = node

	table.length++
}

func (table *Hashtable) Remove(key interface{}) {
	if table == nil {
		panic("A hashtable cannot be nil")
	}

	if key == nil {
		panic("A hashtable key cannot be nil")
	}

	index := table.getBucketIndex(key)

	var prev *LLNode = nil

	node := table.buckets[index]
	for node != nil {
		item := node.Data.(hashItem)
		if item.key == key {
			if prev == nil {
				table.buckets[index] = node.Next
			} else {
				prev.Next = node.Next
			}
			return
		}

		prev = node
		node = node.Next
	}
}

func (table Hashtable) bucketCount() int {
	return len(table.buckets)
}

func (table Hashtable) getBucketIndex(key interface{}) int {
	h := table.hash.calHashCode(key)
	count := table.bucketCount()
	index := h % count
	return index
}
