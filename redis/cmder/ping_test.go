package cmder

import (
	"github.com/smartystreets/goconvey/convey"
	"goredis/redis/reply"
	"testing"
)

func TestPing(t *testing.T) {
	convey.Convey("test ping command", t, func() {
		convey.Convey("ping without args", func() {
			r := Ping(nil, "", nil)
			convey.So(r.Reply(), convey.ShouldResemble, reply.PongReply.Reply())
		})

		convey.Convey("ping with args", func() {
			r := Ping(nil, "", [][]byte{[]byte("hello")})
			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeSimpleStrReply("hello").Reply())
		})
	})
}
