package cmd

import (
	"goredis/inf"
)

const (
	SimpleString = '+'
	Error        = '-'
	BulkString   = '$'
	Array        = '*'
	CR           = '\r'
	LF           = '\n'
	CRLF         = "\r\n"
)

type CmderFunc func(db inf.DB, args [][]byte) Reply
