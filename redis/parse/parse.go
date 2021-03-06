package parse

import (
	"bufio"
	"goredis/inf"
	"strconv"
	"strings"
)

func Array(reader *bufio.Reader) [][]byte {
	sizeStr, _ := reader.ReadString(inf.LF)
	sizeStr = strings.TrimSpace(sizeStr)
	size, _ := strconv.Atoi(sizeStr)
	r := make([][]byte, 0, size)
	for i := 0; i < size; i++ {
		sizeStr, _ := reader.ReadString(inf.LF)
		sizeStr = sizeStr[1:]
		sizeStr = strings.TrimSpace(sizeStr)
		size, _ := strconv.Atoi(sizeStr)
		bs := make([]byte, size)
		reader.Read(bs)
		reader.ReadBytes('\n')
		r = append(r, bs)
	}
	return r
}
