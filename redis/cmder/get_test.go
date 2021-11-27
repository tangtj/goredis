package cmder

import (
	"github.com/smartystreets/goconvey/convey"
	"goredis/inf"
	"goredis/redis/reply"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	convey.Convey("test get command", t, func() {

		client := &inf.Client{
			Db: inf.MakeDb(),
		}

		convey.Convey("get no key", func() {
			err := Get(client, "get", [][]byte{})
			convey.So(err.Reply(), convey.ShouldResemble, reply.ErrArgsNumber("get").Reply())
		})

		convey.Convey("get empty key", func() {
			err := Get(client, "get", [][]byte{[]byte("empty key")})
			convey.So(err.Reply(), convey.ShouldResemble, reply.NilReply.Reply())
		})

		convey.Convey("get key", func() {
			Set(client, "set", [][]byte{[]byte("key"), []byte("value")})
			err := Get(client, "get", [][]byte{[]byte("key")})
			convey.So(err.Reply(), convey.ShouldResemble, reply.MakeBulkReply("value").Reply())
		})

		convey.Convey("get expire key", func() {
			SetEX(client, "setex", [][]byte{[]byte("key"), []byte("10"), []byte("value")})
			err := Get(client, "get", [][]byte{[]byte("key")})
			convey.So(err.Reply(), convey.ShouldResemble, reply.MakeBulkReply("value").Reply())
		})

		convey.Convey("get over expire key", func() {
			SetEX(client, "setex", [][]byte{[]byte("key"), []byte("1"), []byte("value")})
			time.Sleep(time.Second * 2)
			err := Get(client, "get", [][]byte{[]byte("key")})
			convey.So(err.Reply(), convey.ShouldResemble, reply.NilReply.Reply())
		})
	})
}
