package server

import (
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/city-lib/repositories"
	"rac_oblak_proj/config"
	"rac_oblak_proj/interfaces"
	"rac_oblak_proj/models"
)

type CityLibServer struct {
	rentals *repositories.RentalRepo
	books   *repositories.BookRepo
	*baseserver.BaseServer
	loggedInUsers map[string]*models.User
	config        *config.Config
}

func New() interfaces.Server {
	return &CityLibServer{
		loggedInUsers: make(map[string]*models.User),
	}
}
