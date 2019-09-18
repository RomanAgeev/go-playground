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

func (table *Hashtable) Add(key interface{}, data interface{}) {
	if table == nil {
		panic("A hashtable cannot be nil")
	}

	if key == nil {
		panic("A hashtable key cannot be nil")
	}

	bucketIndex := table.getBucketIndex(key)

	bucket := table.buckets[bucketIndex]
	bucket = NewLLNode(bucket, data)
	table.buckets[bucketIndex] = bucket

	table.length++
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
