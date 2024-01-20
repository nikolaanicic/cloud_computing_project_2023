package server

import (
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/interfaces"
)

type CentralLibServer struct {
	*baseserver.BaseServer
}

func New() interfaces.Server {
	return &CentralLibServer{}
}
