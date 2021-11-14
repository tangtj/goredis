package handler

type Option struct {
}

func NewOption() *Option {
	return &Option{}
}

type Apply func(option *Option)
