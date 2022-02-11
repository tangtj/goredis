package cmder

import (
	"github.com/smartystreets/goconvey/convey"
	"goredis/inf"
	"goredis/redis/reply"
	"testing"
)

func TestAppend(t *testing.T) {
	convey.Convey("test append command", t, func() {

		client := &inf.Client{
			Db: inf.MakeDb(),
		}

		convey.Convey("append no exit key", func() {
			r := Append(client, "append", ConvertToArgs("1", "a"))
			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeIntReply(1).Reply())
		})

		convey.Convey("append exit key", func() {
			Set(client, "set", ConvertToArgs("b", "1"))
			r := Append(client, "append", ConvertToArgs("b", "b"))
			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeIntReply(2).Reply())
		})
	})
}
