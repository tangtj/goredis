package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
)

func StrLen(c *inf.Client, command string, args [][]byte) inf.Reply {
	if len(args) < 1 {
		return reply.ErrArgsNumber(command)
	}
	locker := c.Db.GetData().Locker()
	locker.RLock()
	defer locker.RUnlock()

	key := string(args[0])
	d := c.Db.GetData()
	val, has := d.Find(key)
	if !has {
		return reply.MakeIntReply(0)
	}
	v := val.(*inf.DataEntity)
	return reply.MakeIntReply(len(v.Val.(string)))
}
