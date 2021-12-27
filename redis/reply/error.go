package reply

import (
	"fmt"
	"goredis/inf"
	"sync"
)

var ErrNum = MakeErrReply("ERR value is not an integer or out of range")

var _argsErrStr = "ERR wrong number of arguments for '%s' command"

var _unknownErrStr = "ERR unknown command `%s`, with args beginning with:"

var _errArgsMap sync.Map

var _errUnknownMap sync.Map

func ErrArgsNumber(command string) *ErrReply {
	if v, ok := _errArgsMap.Load(command); ok {
		return v.(*ErrReply)
	} else {
		e := MakeErrReply(fmt.Sprintf(_argsErrStr, command))
		_errArgsMap.Store(command, e)
		return e
	}
}

func ErrUnknownMap(command string) inf.Reply {
	if v, ok := _errUnknownMap.Load(command); ok {
		return v.(*ErrReply)
	} else {
		e := MakeErrReply(fmt.Sprintf(_unknownErrStr, command))
		_errUnknownMap.Store(command, e)
		return e
	}
}
