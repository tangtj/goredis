package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
)

func Set(c *inf.Client, command string, args [][]byte) inf.Reply {
	if len(args) <= 1 {
		return reply.MakeErrReply("ERR wrong number of arguments for 'set' command")
	}

	key := string(args[0])
	value := &inf.DataEntity{
		Type: inf.StringType,
		TTl:  0,
		Val:  string(args[1]),
	}
	err := c.Db.GetData().Add(key, value)
	if err != nil {
		return reply.MakeErrReply(err.Error())
	}
	return reply.OKReply
}
