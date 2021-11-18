package inf

const (
	SimpleString = '+'
	Error        = '-'
	BulkString   = '$'
	Array        = '*'
	CR           = '\r'
	LF           = '\n'
	CRLF         = "\r\n"
)

type CmderFunc func(c *Client, args [][]byte) Reply
