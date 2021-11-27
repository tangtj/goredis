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
	_, exit := d.Find(key)
	if exit {
		return reply.MakeIntReply(0)
	}

	value := &inf.DataEntity{
		Type: inf.StringType,
		TTl:  0,
		Val:  string(args[2]),
	}
	d.Add(key, value)
	return reply.MakeIntReply(1)
}
