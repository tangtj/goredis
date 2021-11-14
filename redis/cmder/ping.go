package cmder

import (
	"goredis/inf"
	"goredis/inf/cmd"
	"goredis/redis/reply"
)

func Ping(_ inf.DB, args [][]byte) cmd.Reply {
	if len(args) == 0 {
		return reply.PongReply
	} else {
		return reply.MakeSimpleStrReply(string(args[0]))
	}
}
