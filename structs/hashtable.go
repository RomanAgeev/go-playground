package structs

import (
	"math/rand"
	"time"
)

const initialSize = 31
const prime = 31

type HashCoder interface {
	GetHashCode() int
}

type Hashtable struct {
	buckets []*LLNode
	length  int
	rnd     int
}

type hashItem struct {
	key interface{}
	val interface{}
}

func NewHashtable() *Hashtable {
	randSrc := rand.NewSource(time.Now().Unix())
	randGen := rand.New(randSrc)

	return &Hashtable{
		buckets: make([]*LLNode, initialSize),
		length:  0,
		rnd:     randGen.Int(),
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
	return getHashCode(key, table.rnd) % table.bucketCount()
}

func getHashCode(key interface{}, rnd int) (hash int) {
	switch k := key.(type) {
	case int:
		hash = k * rnd
	case string:
		hash = rnd
		for _, b := range []byte(k) {
			hash = prime*hash + int(b)
		}
	case HashCoder:
		hash = k.GetHashCode()
	default:
		panic("Cannot calc hashcode for the key")
	}
	return
}
