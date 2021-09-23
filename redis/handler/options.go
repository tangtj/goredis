package handler

import (
	"goredis/interface/cmd"
	"sync"
)

type Option struct {
	locker  sync.Locker
	command map[string]cmd.Cmder
}

func NewOption() *Option {
	return &Option{
		locker: &sync.Mutex{},
	}
}

func (o *Option) AddCommand(c cmd.Cmder) *Option {
	o.locker.Lock()
	defer o.locker.Unlock()
	if o.command == nil {
		o.command = make(map[string]cmd.Cmder)
	}
	if c != nil {
		o.command[c.Cmd()] = c
	}
	return o
}

type Apply func(option *Option)
