package cmd

type Reply interface {
	Reply() []byte
}
