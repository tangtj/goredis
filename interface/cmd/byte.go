package cmd

var _ Reply = (*ByteReply)(nil)

type ByteReply struct {
	bs []byte
}

func NewByteReply(bs []byte) *ByteReply {
	return &ByteReply{bs: bs}
}

func (b ByteReply) Reply() []byte {
	return b.bs
}
