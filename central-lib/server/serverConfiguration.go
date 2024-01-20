package server

import (
	"log"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/data_context"
	"rac_oblak_proj/interfaces"
)

func (s *CentralLibServer) Configure(logger *log.Logger, data *data_context.DataContext, host string) interfaces.Server {
	s.BaseServer = baseserver.New(host, logger)

	return s
}
