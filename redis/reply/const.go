package reply

import "goredis/inf/cmd"

var PongReply = (cmd.Reply)(MakeSimpleStrReply("PONG"))

var OKReply = (cmd.Reply)(MakeSimpleStrReply("OK"))

var NilReply = (cmd.Reply)(&bytesReply{bytes: []byte("$-1\r\n")})
