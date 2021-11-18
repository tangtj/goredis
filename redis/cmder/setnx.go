package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
)

func SetNX(c *inf.Client, command string, args [][]byte) inf.Reply {
	locker := c.Db.GetData().Locker()
	locker.Lock()
	defer locker.Unlock()

	d := c.Db.GetData()

	key := string(args[0])
	value := string(args[1])
	_, exit := d.Find(key)
	if exit {
		return &reply.ErrReply{Status: "111"}
	}
	d.Add(key, value)
	return reply.OKReply
}
