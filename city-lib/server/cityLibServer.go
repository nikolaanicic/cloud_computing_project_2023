package server

import (
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/city-lib/repositories"
	"rac_oblak_proj/interfaces"
)

type CityLibServer struct {
	rentals *repositories.RentalRepo
	books   *repositories.BookRepo
	*baseserver.BaseServer
}

func New() interfaces.Server {
	return &CityLibServer{}
}
