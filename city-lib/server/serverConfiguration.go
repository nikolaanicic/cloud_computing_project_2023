package server

import (
	"log"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/city-lib/repositories"
	"rac_oblak_proj/config"
	"rac_oblak_proj/data_context"
	"rac_oblak_proj/interfaces"

	"github.com/go-sql-driver/mysql"
)

func (s *CityLibServer) setBookRepo(books *repositories.BookRepo) {
	s.books = books
}

func (s *CityLibServer) setRentalsRepo(rentals *repositories.RentalRepo) {
	s.rentals = rentals
}

func (s *CityLibServer) setConfiguration(config *config.Config) {
	s.config = config
}

func (s *CityLibServer) Configure(logger *log.Logger, config *config.Config) (interfaces.Server, error) {

	ctx, err := data_context.NewDataContext(mysql.Config{
		User:      config.User,
		Passwd:    config.Password,
		Net:       "tcp",
		Addr:      config.CityDbHost,
		DBName:    config.CityDbName,
		ParseTime: true,
	})

	if err != nil {
		return nil, err
	}

	s.setBookRepo(repositories.NewBookRepo(ctx))
	s.setRentalsRepo(repositories.NewRentalRepo(ctx))
	s.setConfiguration(config)

	s.BaseServer = baseserver.New(config.CityServer, logger, ctx)

	s.RegisterPipelines()

	return s, nil
}

func (s *CityLibServer) RegisterPipelines() {
	for _, p := range s.getPipelines() {
		s.RegisterPipeline(p)
	}
}
