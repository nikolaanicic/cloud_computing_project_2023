package interfaces

import (
	"log"
	"rac_oblak_proj/data_context"
)

type Server interface {
	Serve()
	Configure(logger *log.Logger, data *data_context.DataContext, host string) Server
}
