package dict

import (
	"testing"
)

func TestDict_Find(t *testing.T) {
	dict := NewDict()
	dict.Add("a", 1)

	val, _ := dict.Find("a")
	print(val)

}
