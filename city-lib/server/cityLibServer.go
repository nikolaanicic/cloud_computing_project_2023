package server

import (
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/city-lib/repositories"
	sessionmanager "rac_oblak_proj/city-lib/server/sessionManager"
	"rac_oblak_proj/config"
	"rac_oblak_proj/interfaces"
)

type CityLibServer struct {
	rentals *repositories.RentalRepo
	books   *repositories.BookRepo
	*baseserver.BaseServer
	config     *config.Config
	sessionmgr *sessionmanager.SessionManager
}

func New() interfaces.Server {
	return &CityLibServer{
		sessionmgr: sessionmanager.New(),
	}
}
