package redis

import (
	"goredis/inf"
	"goredis/redis/cmder"
	"goredis/redis/reply"
)

type Server struct {
	DB []*inf.DB

	cmderMap map[string]inf.CmderFunc
}

func (s *Server) Exec(c *inf.Client, command string, args [][]byte) inf.Reply {
	if cmd, ok := s.cmderMap[command]; ok {
		return cmd(c, command, args)
	}
	return reply.MakeErrReply("unSupport command")
}

func MakeServer() *Server {
	s := &Server{}
	dbs := make([]*inf.DB, 16)
	for i := 0; i < len(dbs); i++ {
		dbs[i] = inf.MakeDb()
	}
	s.DB = dbs
	s.cmderMap = cmder.CmdMap
	return s
}
