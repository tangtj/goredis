package cmder

import (
	"github.com/smartystreets/goconvey/convey"
	"goredis/inf"
	"goredis/redis/reply"
	"testing"
)

func TestIncr(t *testing.T) {
	convey.Convey("test incr command", t, func() {

		client := &inf.Client{
			Db: inf.MakeDb(),
		}

		convey.Convey("incr no key", func() {
			r := Incr(client, "incr", ConvertToArgs("a"))
			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeIntReply(1).Reply())
		})

		convey.Convey("incr key has val 1", func() {
			Set(client, "set", ConvertToArgs("b", "1"))
			r := Incr(client, "incr", ConvertToArgs("b"))
			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeIntReply(2).Reply())
		})

		convey.Convey("incr key not number", func() {
			Set(client, "set", ConvertToArgs("c", "-"))
			r := Incr(client, "incr", ConvertToArgs("c"))
			convey.So(r.Reply(), convey.ShouldResemble, reply.ErrNum.Reply())
		})

	})
}
