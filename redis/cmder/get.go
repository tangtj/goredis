package cmder

import (
	"goredis/inf"
	"goredis/inf/cmd"
	"goredis/redis/reply"
)

func Get(db inf.DB, args [][]byte) cmd.Reply {
	if len(args) != 1 {
		return reply.MakeErrReply("error params")
	}
	val, has := db.GetData().Find(string(args[0]))
	// 不存在这个 key
	if !has {
		return reply.NilReply
	}
	v := val.(string)
	return reply.MakeBulkReply(v)
}
