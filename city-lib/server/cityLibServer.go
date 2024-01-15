package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rac_oblak_proj/city-lib/repositories"
	requestmodels "rac_oblak_proj/request_models"
)

type CityLibServer struct {
	rentals *repositories.RentalRepo
	books   *repositories.BookRepo
	logger  *log.Logger
}

func New() *CityLibServer {
	return &CityLibServer{}
}

func (s *CityLibServer) WithLogger(logger *log.Logger) *CityLibServer {
	s.logger = logger
	return s
}

func (s *CityLibServer) WithBookRepo(books *repositories.BookRepo) *CityLibServer {
	s.books = books
	return s
}

func (s *CityLibServer) WithRentalsRepo(rentals *repositories.RentalRepo) *CityLibServer {
	s.rentals = rentals
	return s
}

func (s *CityLibServer) setEncodingHeaders(w http.ResponseWriter) {

	w.Header().Add("Content-Type", "application/json")
}

func (s *CityLibServer) HandleGetAllBooksRequest() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			s.logger.Println(ErrBadReqeustMethod(r.Method))
			http.Error(w, ErrBadRequest(ErrBadReqeustMethod(r.Method)).Error(), http.StatusBadRequest)
			return
		}

		books, err := s.books.GetAll()

		if err != nil {
			s.logger.Println(err)
			http.Error(w, fmt.Sprintf("server error: %v", err), http.StatusInternalServerError)
			return
		}

		s.setEncodingHeaders(w)

		if _, err := w.Write(books.AsJson()); err != nil {
			s.logger.Println(err)
			http.Error(w, fmt.Sprintf("server error: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func (s *CityLibServer) HandleInsertBookRequest() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var insertBookRequest requestmodels.InsertBookRequest

		bodyData, err := io.ReadAll(r.Body)

		if err != nil {
			s.logger.Println(err)
			http.Error(w, fmt.Sprintf("can't read request: %v", err), http.StatusInternalServerError)
			return

		}

		if err := json.Unmarshal(bodyData, &insertBookRequest); err != nil {
			s.logger.Println(err)
			http.Error(w, fmt.Sprintf("Bad Request: %v", err), http.StatusBadRequest)
			return
		}

		s.logger.Println(insertBookRequest)
		result, err := s.books.Insert(insertBookRequest)

		if err != nil {
			s.logger.Println(err)
			http.Error(w, fmt.Sprintf("Bad Request: %v", err), http.StatusBadRequest)
			return
		}

		s.setEncodingHeaders(w)

		if _, err = w.Write(result.AsJson()); err != nil {
			s.logger.Println(err)
			http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			return
		}
	}
}
