package reply

import (
	"goredis/inf"
)

var PongReply = (inf.Reply)(MakeSimpleStrReply("PONG"))

var OKReply = (inf.Reply)(MakeSimpleStrReply("OK"))

var NilReply = (inf.Reply)(&bytesReply{bytes: []byte("$-1\r\n")})
