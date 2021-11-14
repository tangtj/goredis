package dict

import (
	"strings"
	"sync"
)

const (
	// DefaultDictSize is the default size of the dictionary.
	minCapacity = 128

	// DefaultDictSize is the default size of the dictionary.
	maxCapacity = 1 << 16
)

type Dict struct {
	shards     []*Shard
	shardCount int
	count      int
}

type Shard struct {
	table  map[string]interface{}
	locker sync.RWMutex
}

func (dict *Dict) Find(key string) (interface{}, bool) {
	h := genCaseHashFunction(key)
	idx := (dict.shardCount - 1) & h
	shard := dict.shards[idx]
	val, ok := shard.table[key]
	return val, ok
}

// Add is add dict entry
func (dict *Dict) Add(key string, val interface{}) error {
	h := genCaseHashFunction(key)
	idx := (dict.shardCount - 1) & h
	shard := dict.shards[idx]
	shard.table[key] = val
	return nil
}

func genCaseHashFunction(key string) int {
	var hash int = 5381
	lens := len(key)
	chars := []rune(strings.ToLower(key))
	for i := lens - 1; i >= 0; i-- {
		hash = ((hash << 5) + hash) + int(chars[i])
	}
	return hash
}

func NewDict(cap int) *Dict {
	cap = computeCapacity(cap)
	dict := &Dict{
		shardCount: cap,
		shards:     make([]*Shard, cap),
	}
	for i := uint(0); i < uint(cap); i++ {
		dict.shards[i] = &Shard{
			table: make(map[string]interface{}),
		}
	}
	return dict
}

func computeCapacity(param int) int {
	if param <= minCapacity {
		return minCapacity
	}
	n := param - 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	if n < 0 || n >= maxCapacity {
		return maxCapacity
	} else {
		return n + 1
	}
}

/* Thomas Wang's 32 bit Mix Function */
func dictIntHashFunction(key int) uint {
	var hash = uint(key)
	hash = (hash << 0x10) + hash
	hash = hash ^ (hash >> 0x1F)
	hash = (hash << 0x10) + hash
	hash = hash ^ (hash >> 0x1F)
	return hash
}

//func dictIntHashFunction(key uint) {
//key += ~(key << 15)
//	key ^=  (key >> 10)
//	key +=  (key << 3)
//	key ^=  (key >> 6)
//	key += ~(key << 11)
//	key ^=  (key >> 16)
//	return key
//}
