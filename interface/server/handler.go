package server

import (
	"context"
	"net"
)

type HandleFunc func(ctx context.Context, conn net.Conn)

type Handler interface {
	Close() error
	Handle(ctx context.Context, conn net.Conn)
}
