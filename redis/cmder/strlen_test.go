package cmder

import (
	"github.com/smartystreets/goconvey/convey"
	"goredis/inf"
	"goredis/redis/reply"
	"testing"
)

func TestStrLen(t *testing.T) {

	convey.Convey("test command strlen", t, func() {

		c := &inf.Client{
			Db: inf.MakeDb(),
		}

		convey.Convey("test err args strlen", func() {
			r := StrLen(c, "strlen", ConvertToArgs())

			convey.So(r.Reply(), convey.ShouldResemble, reply.ErrArgsNumber("strlen").Reply())
		})

		convey.Convey("test str strlen", func() {
			Set(c, "set", ConvertToArgs("a", "1"))

			r := StrLen(c, "strlen", ConvertToArgs("a"))

			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeIntReply(1).Reply())

			Set(c, "set", ConvertToArgs("a", "1xxxxxxxxxxxxxxxxxxxxxxxx"))

			r = StrLen(c, "strlen", ConvertToArgs("a"))

			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeIntReply(len("1xxxxxxxxxxxxxxxxxxxxxxxx")).Reply())
		})

		convey.Convey("test not exsit key strlen", func() {
			r := StrLen(c, "strlen", ConvertToArgs("a"))

			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeIntReply(0).Reply())
		})

	})
}
