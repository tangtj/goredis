package cmd

import "goredis/interface/cmd"

var _ cmd.Cmder = (*PingCmder)(nil)

type PingCmder struct {
}

func (p PingCmder) Cmd() string {
	return "ping"
}

func (p *PingCmder) Exec(args [][]byte) cmd.Reply {
	if len(args) == 0 {
		return cmd.NewByteReply([]byte("pong"))
	} else {
		return cmd.NewByteReply(args[0])
	}
}
