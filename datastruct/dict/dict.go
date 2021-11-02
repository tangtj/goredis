package dict

import (
	"goredis/datastruct/list"
	"strings"
)

type Dict struct {
	rehashIdx int
	dictHt    [2]list.List
}

type DictEntry struct {
	key   string
	value interface{}
}

// Key get dict entry key
func (e *DictEntry) Key() string {
	return e.key
}

// Value get dict entry value
func (e *DictEntry) Value() interface{} {
	return e.value
}

func (dict *Dict) Find(key string) (interface{}, error) {
	h := genCaseHashFunction(key)

	for tableIdx := 0; tableIdx <= 1; tableIdx++ {
		idx := h&dict.dictHt[tableIdx].GetLen() - 1

		he := dict.dictHt[tableIdx].GetIdx(int(idx)).(*list.List)

		node := he.GetNode(0)
		for he != nil {

			t := node.GetValue().(*DictEntry)

			if t.Key() == key {
				return t.value, nil
			}
			node = node.GetNext()
		}

	}
	return nil
}

func genCaseHashFunction(key string) uint {
	var hash uint = 5381
	lens := len(key)
	chars := []rune(strings.ToLower(key))
	for i := lens - 1; i > 0; i-- {
		hash = ((hash << 5) + hash) + uint(chars[i])
	}
	return hash
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
