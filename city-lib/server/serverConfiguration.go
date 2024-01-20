package server

import (
	"log"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/city-lib/repositories"
	"rac_oblak_proj/data_context"
	"rac_oblak_proj/interfaces"
)

func (s *CityLibServer) setBookRepo(books *repositories.BookRepo) {
	s.books = books
}

func (s *CityLibServer) setRentalsRepo(rentals *repositories.RentalRepo) {
	s.rentals = rentals
}

func (s *CityLibServer) Configure(logger *log.Logger, data *data_context.DataContext, host string) interfaces.Server {
	s.setBookRepo(repositories.NewBookRepo(data))
	s.setRentalsRepo(repositories.NewRentalRepo(data))

	s.BaseServer = baseserver.New(host, logger)

	s.BaseServer.RegisterHandler("/books/getAll", s.handleGetAllBooksRequest)
	s.BaseServer.RegisterHandler("/books/insert", s.handleInsertBookRequest)

	return s
}
