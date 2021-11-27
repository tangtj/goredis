package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
	"strconv"
	"time"
)

func SetEX(c *inf.Client, cmd string, args [][]byte) inf.Reply {
	d := c.Db.GetData()
	expire := c.Db.GetExpire()

	// 判断是否有效
	if len(args) != 3 {
		return reply.ErrArgsNumber(cmd)
	}

	key := string(args[0])

	s := string(args[1])
	second, err := strconv.Atoi(s)
	if err != nil {
		return reply.ErrNum
	}
	now := time.Now().Add(time.Duration(second) * time.Second).UnixMilli()

	expire.Add(key, now)

	value := &inf.DataEntity{
		Type: inf.StringType,
		TTl:  0,
		Val:  string(args[2]),
	}
	d.Add(key, value)
	return reply.OKReply
}
