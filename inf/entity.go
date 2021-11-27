package inf

type DataType int

const (
	StringType DataType = iota
)

type DataEntity struct {
	Type DataType
	TTl  int
	Val  interface{}
}
