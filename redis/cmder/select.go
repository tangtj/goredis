package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
	"strconv"
)

func Select(c *inf.Client, _ string, args [][]byte) inf.Reply {
	s := string(args[0])
	if len(s) <= 0 {
		return reply.MakeErrReply("ERR wrong number of arguments for 'select' command")
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return reply.MakeErrReply("ERR value is not an integer or out of range")
	}
	info := c.Server.GetInfo()

	// db 下标是从 0 开始的
	if num >= info.DbNum || num < 0 {
		return reply.MakeErrReply("ERR DB index is out of range")
	}
	c.Db = c.Server.GetDB()[num]
	return reply.OKReply
}
