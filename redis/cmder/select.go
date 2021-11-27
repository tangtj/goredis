package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
	"strconv"
)

func Select(c *inf.Client, cmd string, args [][]byte) inf.Reply {
	s := string(args[0])
	if len(s) <= 0 {
		return reply.ErrArgsNumber(cmd)
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return reply.ErrNum
	}
	info := c.Server.GetInfo()

	// db 下标是从 0 开始的
	if num >= info.DbNum || num < 0 {
		return reply.MakeErrReply("ERR DB index is out of range")
	}
	c.Db = c.Server.GetDB()[num]
	return reply.OKReply
}
