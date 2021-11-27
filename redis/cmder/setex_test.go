package cmder

import (
	"github.com/smartystreets/goconvey/convey"
	"goredis/inf"
	"goredis/redis/reply"
	"testing"
	"time"
)

func TestSetEX(t *testing.T) {

	convey.Convey("test command setex", t, func() {

		command := "setex"

		client := &inf.Client{
			Db: inf.MakeDb(),
		}

		now := time.Now().UnixMilli()

		args := [][]byte{
			[]byte("key"),
			[]byte("10"),
			[]byte("value"),
		}
		r := SetEX(client, command, args)

		d := client.Db.GetData()
		e := client.Db.GetExpire()

		convey.Convey("test setex ok reply", func() {
			convey.So(r.Reply(), convey.ShouldResemble, reply.OKReply.Reply())
		})

		convey.Convey("test setex err time", func() {
			r := SetEX(client, command, [][]byte{
				[]byte("key"),
				[]byte("aa"),
				[]byte("value"),
			})
			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeErrReply("ERR value is not an integer or out of range").Reply())
		})

		convey.Convey("test setex err args", func() {
			r := SetEX(client, command, [][]byte{
				[]byte("key"),
				[]byte("aa"),
				[]byte("value"),
				[]byte(" "),
			})
			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeErrReply("ERR wrong number of arguments for 'setex' command").Reply())
		})

		convey.Convey("test db value", func() {
			value, _ := d.Find("key")
			convey.So(value.(string), convey.ShouldEqual, "value")
		})

		convey.Convey("test expire value", func() {
			value, _ := e.Find("key")
			convey.So(value.(int64), convey.ShouldBeGreaterThan, 10+now)
		})

	})
}
