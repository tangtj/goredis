package cmder

import (
	"goredis/inf"
	"goredis/redis/reply"
	"strings"
)

func Append(c *inf.Client, cmd string, args [][]byte) inf.Reply {
	l := len(args)
	if l < 2 {
		return reply.ErrArgsNumber(cmd)
	}

	key := string(args[0])
	val, has := c.Db.GetData().Find(key)
	size := 0
	// 不存在这个 key
	if has {
		v := val.(*inf.DataEntity)
		newStr := strings.Join([]string{v.Val.(string), string(args[1])}, "")
		v.Val, size = newStr, len(newStr)
	} else {
		key := string(args[0])
		value := &inf.DataEntity{
			Type: inf.StringType,
			TTl:  0,
			Val:  string(args[1]),
		}
		err := c.Db.GetData().Add(key, value)
		if err != nil {
			return reply.MakeErrReply(err.Error())
		}
		size = len(args[1])
	}
	return reply.MakeIntReply(size)
}
