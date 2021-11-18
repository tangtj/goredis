package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
)

func Set(c *inf.Client, command string, args [][]byte) inf.Reply {
	if len(args) < 1 {
		return reply.MakeErrReply("err redis")
	}
	err := c.Db.GetData().Add(string(args[0]), string(args[1]))
	if err != nil {
		return reply.MakeErrReply(err.Error())
	}
	return reply.MakeIntReply(1)
}
