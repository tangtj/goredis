package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
)

func Ping(_ *inf.Client, args [][]byte) inf.Reply {
	if len(args) == 0 {
		return reply.PongReply
	} else {
		return reply.MakeSimpleStrReply(string(args[0]))
	}
}
