package cmd

const (
	SimpleString = '+'
	Error        = '-'
	BulkString   = '$'
	Array        = '*'
	CR           = '\r'
	LF           = '\n'
)

type Cmder interface {
	Cmd() string
	Exec(args [][]byte) Reply
}
