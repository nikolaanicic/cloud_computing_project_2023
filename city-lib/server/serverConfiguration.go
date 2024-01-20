package server

import (
	"log"
	"rac_oblak_proj/city-lib/repositories"
	"rac_oblak_proj/data_context"
	"rac_oblak_proj/interfaces"
)

func (s *CityLibServer) setHost(host string) {
	s.addr = host
}

func (s *CityLibServer) setLogger(logger *log.Logger) {
	s.logger = logger
}

func (s *CityLibServer) setBookRepo(books *repositories.BookRepo) {
	s.books = books
}

func (s *CityLibServer) setRentalsRepo(rentals *repositories.RentalRepo) {
	s.rentals = rentals
}

func (s *CityLibServer) Configure(logger *log.Logger, data *data_context.DataContext, host string) interfaces.Server {
	s.setLogger(logger)
	s.setBookRepo(repositories.NewBookRepo(data))
	s.setRentalsRepo(repositories.NewRentalRepo(data))
	s.setHost(host)

	return s

}
