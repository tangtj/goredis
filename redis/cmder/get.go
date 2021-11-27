package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
	"time"
)

func Get(c *inf.Client, cmd string, args [][]byte) inf.Reply {
	if len(args) != 1 {
		return reply.ErrArgsNumber(cmd)
	}

	key := string(args[0])
	val, has := c.Db.GetData().Find(key)
	// 不存在这个 key
	if !has {
		return reply.NilReply
	}

	ts, h := c.Db.GetExpire().Find(key)
	now := time.Now().UnixMilli()
	if h && now > ts.(int64) {
		return reply.NilReply
	}
	v := val.(*inf.DataEntity)
	return reply.MakeBulkReply(v.Val.(string))
}
