package interfaces

import (
	"log"
	"rac_oblak_proj/config"
)

type Server interface {
	Serve()
	Configure(logger *log.Logger, config *config.Config) (Server, error)
	RegisterPipelines()
}
