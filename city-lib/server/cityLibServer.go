package server

import (
	"log"
	"net/http"
	"rac_oblak_proj/city-lib/repositories"
	"rac_oblak_proj/errors/http_errors"
	"rac_oblak_proj/interfaces"
)

var encoding = "application/json"

type CityLibServer struct {
	rentals *repositories.RentalRepo
	books   *repositories.BookRepo
	logger  *log.Logger
	mux     *http.ServeMux
	addr    string

	handlers map[string]func(http.ResponseWriter, *http.Request) *http_errors.HttpErrorResponse
}

func New() interfaces.Server {
	return &CityLibServer{
		mux:      http.NewServeMux(),
		handlers: make(map[string]func(http.ResponseWriter, *http.Request) *http_errors.HttpErrorResponse),
	}
}

func (s *CityLibServer) Serve() {

	s.registerHandlers()

	s.logger.Println("listening on", s.addr)
	if err := http.ListenAndServe(s.addr, s.mux); err != nil {
		s.logger.Fatal(err)
	}
}
