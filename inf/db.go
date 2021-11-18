package inf

import (
	"goredis/datastruct/dict"
)

type DB struct {
	Data    *dict.Dict
	Expires *dict.Dict
}

// MakeDb 创建一个Db
func MakeDb() *DB {
	d := &DB{}

	size := 16
	d.Data = dict.NewDict(size)
	d.Expires = dict.NewDict(size)
	return d
}

func (db *DB) GetData() *dict.Dict {
	return db.Data
}

// GetExpires 获取过期数据
func (db *DB) GetExpire() *dict.Dict {
	return db.Expires
}
