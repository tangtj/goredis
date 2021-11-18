package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
)

func Get(c *inf.Client, args [][]byte) inf.Reply {
	if len(args) != 1 {
		return reply.MakeErrReply("error params")
	}
	val, has := c.Db.GetData().Find(string(args[0]))
	// 不存在这个 key
	if !has {
		return reply.NilReply
	}
	v := val.(string)
	return reply.MakeBulkReply(v)
}
