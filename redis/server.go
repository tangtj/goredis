package redis

import (
	"goredis/inf"
	"goredis/redis/cmder"
	"goredis/redis/reply"
	"strings"
)

type Server struct {
	DB []*inf.DB

	cmderMap map[string]inf.CmderFunc
}

func (s *Server) Exec(c *inf.Client, command string, args [][]byte) inf.Reply {
	command = strings.ToLower(command)
	if cmd, ok := s.cmderMap[command]; ok {
		return cmd(c, command, args)
	}
	return reply.MakeErrReply("unSupport command")
}

func (s *Server) GetInfo() inf.ServerInfo {
	return inf.ServerInfo{
		DbNum: len(s.DB),
	}
}

func (s *Server) GetDB() []*inf.DB {
	return s.DB
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
