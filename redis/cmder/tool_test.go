package cmder

func ConvertToArgs(args ...string) [][]byte {
	if len(args) <= 0 {
		return [][]byte{}
	}
	bytes := make([][]byte, 0, len(args))
	for _, arg := range args {
		bytes = append(bytes, []byte(arg))
	}
	return bytes
}
