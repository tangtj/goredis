package inf

type Server interface {
	GetInfo() ServerInfo
	GetDB() []*DB
}

type ServerInfo struct {
	DbNum int
}
