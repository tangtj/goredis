package reply

import "strconv"
import "goredis/inf/cmd"

type ErrReply struct {
	Status string
}

func (r *ErrReply) Reply() []byte {
	return []byte("-" + r.Status + cmd.CRLF)
}

func MakeErrReply(err string) *ErrReply {
	return &ErrReply{Status: err}
}

type SimpleStrReply struct {
	string string
}

func (r *SimpleStrReply) Reply() []byte {
	return []byte("+" + r.string + cmd.CRLF)
}

func MakeSimpleStrReply(str string) *SimpleStrReply {
	return &SimpleStrReply{string: str}
}

type IntReply struct {
	number int
}

func (r *IntReply) Reply() []byte {
	return []byte(":" + strconv.Itoa(r.number) + cmd.CRLF)
}

type BulkReply struct {
	string string
}

func (r *BulkReply) Reply() []byte {
	return []byte("$" + strconv.Itoa(len(r.string)) + cmd.CRLF + r.string + cmd.CRLF)
}

type bytesReply struct {
	bytes []byte
}

func (b *bytesReply) Reply() []byte {
	return b.bytes
}

func MakeBulkReply(str string) *BulkReply {
	return &BulkReply{string: str}
}
