package server

import (
	"log"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/central-lib/repositories"
	"rac_oblak_proj/config"
	"rac_oblak_proj/data_context"
	"rac_oblak_proj/interfaces"

	"github.com/go-sql-driver/mysql"
)

func (s *CentralLibServer) setUserRepo(user *repositories.UserRepo) {
	s.userRepo = user
}

func (s *CentralLibServer) setConfiguration(config *config.Config) {
	s.cfg = config
}

func (s *CentralLibServer) Configure(logger *log.Logger, config *config.Config) (interfaces.Server, error) {
	ctx, err := data_context.NewDataContext(mysql.Config{
		User:      config.User,
		Passwd:    config.Password,
		Net:       "tcp",
		Addr:      config.CentralDbHost,
		DBName:    config.CentralDbName,
		ParseTime: true,
	})

	if err != nil {
		return nil, err
	}

	s.setUserRepo(repositories.NewUserRepo(ctx))
	s.setConfiguration(config)

	s.BaseServer = baseserver.New(config.CentralServerHost, logger, ctx)

	s.RegisterPipelines()

	return s, nil
}

func (s *CentralLibServer) RegisterPipelines() {
	for _, p := range s.getPipelines() {
		s.BaseServer.RegisterPipeline(p)
	}
}
