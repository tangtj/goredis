package main

import (
	"goredis/redis"
	"goredis/server"
)

func main() {
	cfg := server.Config{Address: ":16379"}
	h := redis.NewHandler()
	server.ListenAndHandleServe(cfg, h)
}
