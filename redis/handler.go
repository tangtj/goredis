package redis

import (
	"bufio"
	"context"
	"goredis/inf"
	"goredis/redis/handler"
	"goredis/redis/parse"
	"log"
	"net"
)

var _ inf.Handler = (*Handler)(nil)

type Handler struct {
	server *Server
}

func NewHandler(options ...handler.Apply) inf.Handler {
	h := Handler{}
	h.server = MakeServer()
	o := handler.NewOption()
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
	client := &inf.Client{
		Db: h.server.DB[0],
	}
	for {
		b, err := reader.ReadByte()
		if err != nil {
			log.Fatalf("read command err : %s", err)
			return
		}
		switch b {
		case inf.Array:
			// *1\r\n$4\r\nping\r\n
			r := parse.Array(reader)

			c := string(r[0])
			reply := h.server.Exec(client, c, r[1:])
			conn.Write(reply.Reply())
		}
	}

}
