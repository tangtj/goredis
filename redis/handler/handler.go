package handler

import (
	"bufio"
	"context"
	"goredis/interface/cmd"
	"goredis/interface/server"
	"goredis/redis/parse"
	"log"
	"net"
)

var _ server.Handler = (*Handler)(nil)

type Handler struct {
	options map[string]cmd.Cmder
}

func NewHandler(options ...Apply) server.Handler {
	h := Handler{}
	o := NewOption()
	for _, f := range options {
		f(o)
	}
	h.options = o.command
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
			reply := h.exec(c, r[1:])
			conn.Write(reply.Reply())
		}
	}

}

func (h *Handler) exec(c string, args [][]byte) cmd.Reply {
	if c, ok := h.options[c]; ok {
		return c.Exec(args)
	}
	return cmd.NewByteReply([]byte("unSupport command"))
}
