package cmder

import (
	"github.com/smartystreets/goconvey/convey"
	"goredis/inf"
	"goredis/redis/reply"
	"testing"
)

type Server struct {
	DB []*inf.DB
}

func (s *Server) GetInfo() inf.ServerInfo {
	return inf.ServerInfo{
		DbNum: len(s.DB),
	}
}

func (s *Server) GetDB() []*inf.DB {
	return s.DB
}

func TestSelect(t *testing.T) {
	command := "select"

	s := &Server{}
	d := make([]*inf.DB, 32)
	for i := 0; i < 32; i++ {
		d[i] = inf.MakeDb()
	}
	s.DB = d

	client := &inf.Client{
		Db:     inf.MakeDb(),
		Server: s,
	}

	convey.Convey("test select db", t, func() {

		convey.Convey("select 1 ok", func() {
			r := Select(client, command, [][]byte{
				[]byte("1"),
			})
			convey.So(r, convey.ShouldResemble, reply.OKReply)
			convey.So(client.Db, convey.ShouldEqual, client.Server.GetDB()[1])
		})

		convey.Convey("select args empty", func() {
			r := Select(client, command, [][]byte{
				[]byte(""),
			})
			convey.So(r, convey.ShouldResemble, reply.MakeErrReply("ERR wrong number of arguments for 'select' command"))
		})
		convey.Convey("select too many args", func() {
			r := Select(client, command, [][]byte{
				[]byte(""),
				[]byte("123123"),
			})
			convey.So(r, convey.ShouldResemble, reply.MakeErrReply("ERR wrong number of arguments for 'select' command"))
		})
		convey.Convey("select outer range db num", func() {
			r := Select(client, command, [][]byte{
				[]byte("80"),
			})
			convey.So(r, convey.ShouldResemble, reply.MakeErrReply("ERR DB index is out of range"))
		})
		convey.Convey("select err db num", func() {
			r := Select(client, command, [][]byte{
				[]byte("-100"),
			})
			convey.So(r, convey.ShouldResemble, reply.MakeErrReply("ERR DB index is out of range"))
		})
		convey.Convey("select err num", func() {
			r := Select(client, command, [][]byte{
				[]byte("a"),
			})
			convey.So(r, convey.ShouldResemble, reply.MakeErrReply("ERR value is not an integer or out of range"))
		})
	})
}
