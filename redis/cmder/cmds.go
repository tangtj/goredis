package cmder

import "goredis/inf"

var CmdMap = map[string]inf.CmderFunc{
	"ping":  Ping,
	"get":   Get,
	"set":   Set,
	"setnx": SetNX,
}
