package cmder

import (
	"github.com/smartystreets/goconvey/convey"
	"goredis/inf"
	"goredis/redis/reply"
	"testing"
)

func TestSetEX(t *testing.T) {

	convey.Convey("test command setex", t, func() {

		command := "setex"

		client := &inf.Client{
			Db: inf.MakeDb(),
		}

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

		convey.Convey("test db value", func() {
			value, _ := d.Find("key")
			convey.So(value.(string), convey.ShouldEqual, "value")
		})

		convey.Convey("test expire value", func() {
			value, _ := e.Find("key")
			convey.So(value.(int), convey.ShouldEqual, 10)
		})

	})
}
