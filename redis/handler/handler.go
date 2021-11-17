package handler

import (
	"bufio"
	"context"
	"goredis/inf/cmd"
	"goredis/inf/server"
	"goredis/redis"
	"goredis/redis/parse"
	"log"
	"net"
)

var _ server.Handler = (*Handler)(nil)

type Handler struct {
	Db *redis.Db
}

func NewHandler(options ...Apply) server.Handler {
	h := Handler{}
	h.Db = redis.MakeDb()
	o := NewOption()
	for _, f := range options {
		f(o)
	}
	return &h
}

func (h *Handler) Close() error {
	return nil
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			log.Fatalf("read command err : %s", err)
			return
		}
		switch b {
		case cmd.Array:
			// *1\r\n$4\r\nping\r\n
			r := parse.Array(reader)

			c := string(r[0])
			reply := h.Db.Exec(c, r[1:])
			conn.Write(reply.Reply())
		}
	}

}
