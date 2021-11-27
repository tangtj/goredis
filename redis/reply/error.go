package reply

import (
	"fmt"
	"sync"
)

var ErrNum = MakeErrReply("ERR value is not an integer or out of range")

var _argsErrStr = "ERR wrong number of arguments for '%s' command"

var _errMap sync.Map

func ErrArgsNumber(command string) *ErrReply {
	if v, ok := _errMap.Load(command); ok {
		return v.(*ErrReply)
	} else {
		e := MakeErrReply(fmt.Sprintf(_argsErrStr, command))
		_errMap.Store(command, e)
		return e
	}
}
