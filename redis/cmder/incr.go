package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
	"strconv"
)

func Incr(c *inf.Client, cmd string, args [][]byte) inf.Reply {
	argsNum := len(args)

	if argsNum != 1 {
		return reply.ErrArgsNumber(cmd)
	}

	key := string(args[0])
	db := c.Db

	data := db.GetData()

	val, has := data.Find(key)
	if !has {
		data.Locker().Lock()
		defer data.Locker().Unlock()

		data.Add(key, &inf.DataEntity{
			Type: inf.StringType,
			TTl:  0,
			Val:  "1",
		})
		return reply.MakeIntReply(1)
	}

	data.Locker().Lock()
	defer data.Locker().Unlock()
	entity := val.(*inf.DataEntity)
	v := entity.Val.(string)

	num, err := strconv.Atoi(v)
	if err != nil {
		return reply.ErrNum
	} else {
		entity.Val = strconv.Itoa(num + 1)
	}
	return reply.MakeIntReply(num + 1)
}
