package inf

type Client struct {
	Id string
	Db *DB

	// 数据库信息
	Server Server

	// 最后活跃时间
	LastActive int64
}
