package server

import (
	"context"
	"goredis/interface/server"
	"log"
	"net"
)

type Config struct {
	Address string `yaml:"address"`
}

func ListenAndHandleServe(cfg Config, handler server.Handler) error {
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Panicf("listen addr : %s , err : %s", cfg.Address, err)
	}

	defer listener.Close()
	defer handler.Close()
	ctx, _ := context.WithCancel(context.Background())
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handler.Handle(ctx, conn)
	}

}
