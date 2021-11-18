package inf

type Client struct {
	Id string
	Db DB

	// 最后活跃时间
	LastActive int64
}
