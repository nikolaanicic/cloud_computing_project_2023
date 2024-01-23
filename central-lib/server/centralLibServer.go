package server

import (
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/central-lib/repositories"
	"rac_oblak_proj/config"
	"rac_oblak_proj/interfaces"
)

type CentralLibServer struct {
	*baseserver.BaseServer
	userRepo *repositories.UserRepo
	cfg      *config.Config
}

func New() interfaces.Server {
	return &CentralLibServer{}
}
