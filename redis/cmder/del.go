package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
)

func Del(c *inf.Client, cmd string, args [][]byte) inf.Reply {
	l := len(args)
	if l <= 0 {
		return reply.ErrArgsNumber(cmd)
	}
	data := c.Db.GetData()
	r := 0
	for i := 0; i < l; i++ {
		key := string(args[i])
		if data.Del(key) {
			r++
		}
	}
	return reply.MakeIntReply(r)
}
