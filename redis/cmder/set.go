package cmder

import (
	"goredis/inf"
	"goredis/inf/cmd"
	"goredis/redis/reply"
)

func Set(db inf.DB, args [][]byte) cmd.Reply {
	if len(args) < 1 {
		return reply.MakeErrReply("err redis")
	}
	err := db.GetData().Add(string(args[0]), string(args[1]))
	if err != nil {
		return reply.MakeErrReply(err.Error())
	}
	return reply.OKReply
}
