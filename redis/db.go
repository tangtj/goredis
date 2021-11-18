package redis

import (
	"goredis/datastruct/dict"
)

type Db struct {
	Data *dict.Dict
}

// MakeDb 创建一个Db
func MakeDb() *Db {
	d := &Db{}
	d.Data = dict.NewDict(16)
	return d
}

func (db *Db) GetData() *dict.Dict {
	return db.Data
}
