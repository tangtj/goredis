package inf

import "goredis/datastruct/dict"

type DB interface {
	GetData() *dict.Dict
}
