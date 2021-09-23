package main

import (
	"goredis/redis/cmd"
	"goredis/redis/handler"
	"goredis/server"
)

func main() {
	cfg := server.Config{Address: ":16379"}
	h := handler.NewHandler(cmd.Register)
	server.ListenAndHandleServe(cfg, h)
}
