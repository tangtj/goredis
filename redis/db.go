package redis

import (
	"goredis/datastruct/dict"
	"goredis/inf/cmd"
	"goredis/redis/cmder"
	"goredis/redis/reply"
)

type Db struct {
	cmdMap map[string]cmd.CmderFunc
	Data   *dict.Dict
}

var _cmdMap = map[string]cmd.CmderFunc{
	"ping": cmder.Ping,
	"get":  cmder.Get,
}

// MakeDb 创建一个Db
func MakeDb() *Db {
	return &Db{
		cmdMap: _cmdMap,
	}
}

func (db *Db) Exec(c string, args [][]byte) cmd.Reply {
	if c, ok := db.cmdMap[c]; ok {
		return c(db, args)
	}
	return reply.MakeErrReply("unSupport command")
}

func (db *Db) GetData() *dict.Dict {
	return db.Data
}
