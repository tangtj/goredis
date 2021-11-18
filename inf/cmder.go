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

type CmderFunc func(c *Client, command string, args [][]byte) Reply
