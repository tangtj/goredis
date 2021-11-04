package dict

import (
	"goredis/datastruct/list"
	"strings"
)

type Dict struct {
	rehashIdx int
	dictHt    [2]*list.List
}

type Entry struct {
	key   string
	value interface{}
}

// Key get dict entry key
func (e *Entry) Key() string {
	return e.key
}

// Value get dict entry value
func (e *Entry) Value() interface{} {
	return e.value
}

func (dict *Dict) Find(key string) (interface{}, error) {
	h := genCaseHashFunction(key)

	for tableIdx := 0; tableIdx <= 1; tableIdx++ {
		idx := h&dict.dictHt[tableIdx].GetLen() - 1

		he := dict.dictHt[tableIdx].GetIdx(int(idx)).(*list.List)

		node := he.GetNode(0)
		for he != nil {

			t := node.GetValue().(*Entry)

			if t.Key() == key {
				return t.value, nil
			}
			node = node.GetNext()
		}

	}
	return nil, nil
}

// Add is add dict entry
func (dict *Dict) Add(key string, val interface{}) error {
	h := genCaseHashFunction(key)

	for tableIdx := 0; tableIdx <= 1; tableIdx++ {
		idx := h & (dict.dictHt[tableIdx].GetLen() - 1)

		he := dict.dictHt[tableIdx].GetIdx(int(idx)).(*list.List)

		node := he.GetNode(0)
		for he != nil {

			t := node.GetValue().(*Entry)

			if t.Key() == key {
				t.value = val
				return nil
			}
			node = node.GetNext()
		}

	}

	dict.dictHt[dict.rehashIdx].AddNodeHead(NewDictEntry(key, val))

	dict.rehashIdx = 1 - dict.rehashIdx

	return nil
}

func genCaseHashFunction(key string) uint {
	var hash uint = 5381
	lens := len(key)
	chars := []rune(strings.ToLower(key))
	for i := lens - 1; i >= 0; i-- {
		hash = ((hash << 5) + hash) + uint(chars[i])
	}
	return hash
}

func NewDictEntry(key string, val interface{}) *Entry {
	return &Entry{key: key, value: val}
}

func NewDict() *Dict {
	return &Dict{
		rehashIdx: 0,
		dictHt: [2]*list.List{
			list.NewList(),
			list.NewList(),
		},
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
