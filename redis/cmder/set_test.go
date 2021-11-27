package cmder

import (
	"github.com/smartystreets/goconvey/convey"
	"goredis/inf"
	"goredis/redis/reply"
	"testing"
)

func TestSet(t *testing.T) {
	convey.Convey("test set command", t, func() {
		command := "set"

		client := &inf.Client{
			Db: inf.MakeDb(),
		}
		convey.Convey("set key 1", func() {
			args := [][]byte{[]byte("key"), []byte("1")}
			r := Set(client, command, args)
			convey.So(r.Reply(), convey.ShouldResemble, reply.OKReply.Reply())
			v := Get(client, "get", [][]byte{[]byte("key")})
			convey.So(v.Reply(), convey.ShouldResemble, reply.MakeBulkReply("1").Reply())
		})

		convey.Convey("set key", func() {
			args := [][]byte{[]byte("key")}
			r := Set(client, command, args)
			convey.So(r.Reply(), convey.ShouldResemble, reply.MakeErrReply("ERR wrong number of arguments for 'set' command").Reply())
		})
	})
}
