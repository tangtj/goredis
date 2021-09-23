package cmd

import (
	"goredis/interface/cmd"
	"goredis/redis/handler"
)

func RegisterCmder(o *handler.Option, c cmd.Cmder) {
	o.AddCommand(c)
}

func Register(o *handler.Option) {
	RegisterCmder(o, &PingCmder{})
}
