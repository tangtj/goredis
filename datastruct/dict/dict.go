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
	locker     sync.RWMutex
}

type Shard struct {
	table  map[string]interface{}
	locker sync.RWMutex
}

func (dict *Dict) Locker() *sync.RWMutex {
	return &dict.locker
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

func (dict *Dict) Del(key string) bool {
	h := genCaseHashFunction(key)
	idx := (dict.shardCount - 1) & h
	shard := dict.shards[idx]
	_, ok := shard.table[key]
	if !ok {
		return false
	}
	delete(shard.table, key)
	return true
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
