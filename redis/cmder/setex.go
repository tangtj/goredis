package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
	"strconv"
	"time"
)

func SetEX(c *inf.Client, _ string, args [][]byte) inf.Reply {
	d := c.Db.GetData()
	expire := c.Db.GetExpire()

	// 判断是否有效
	if len(args) != 3 {
		return reply.MakeErrReply("ERR wrong number of arguments for 'setex' command")
	}

	key := string(args[0])

	s := string(args[1])
	second, err := strconv.Atoi(s)
	if err != nil {
		return reply.MakeErrReply("ERR value is not an integer or out of range")
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
