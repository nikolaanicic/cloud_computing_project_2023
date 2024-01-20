package server

import (
	"log"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/central-lib/repositories"
	"rac_oblak_proj/data_context"
	"rac_oblak_proj/interfaces"
)

func (s *CentralLibServer) setUserRepo(user *repositories.UserRepo) {
	s.userRepo = user
}

func (s *CentralLibServer) Configure(logger *log.Logger, data *data_context.DataContext, host string) interfaces.Server {
	s.setUserRepo(repositories.NewUserRepo(data))

	s.BaseServer = baseserver.New(host, logger)

	s.BaseServer.RegisterHandler("/users/insert", s.handleInsertUser)

	return s
}
