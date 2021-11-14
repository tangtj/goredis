package main

import (
	"goredis/redis/handler"
	"goredis/server"
)

func main() {
	cfg := server.Config{Address: ":16379"}
	h := handler.NewHandler()
	server.ListenAndHandleServe(cfg, h)
}
