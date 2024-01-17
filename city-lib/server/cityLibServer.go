package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"rac_oblak_proj/city-lib/repositories"
	requestmodels "rac_oblak_proj/request_models"
)

var encoding = "application/json"

type CityLibServer struct {
	rentals *repositories.RentalRepo
	books   *repositories.BookRepo
	logger  *log.Logger
	mux     *http.ServeMux
	addr    string

	handlers map[string]func(http.ResponseWriter, *http.Request) *HttpErrorResponse
}

func New() *CityLibServer {
	return &CityLibServer{
		mux:      http.NewServeMux(),
		handlers: make(map[string]func(http.ResponseWriter, *http.Request) *HttpErrorResponse),
	}
}

func (s *CityLibServer) registerHandlers() {
	s.mux.HandleFunc("/", s.rootHandler)

	s.handlers["/books/getAll"] = s.handleGetAllBooksRequest
	s.handlers["/books/insert"] = s.handleInsertBookRequest

}

func (s *CityLibServer) Serve() {

	s.registerHandlers()

	s.logger.Println("listening on", s.addr)
	if err := http.ListenAndServe(s.addr, s.mux); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *CityLibServer) WithHost(host string) *CityLibServer {
	s.addr = host
	return s
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
	w.Header().Add("Content-Type", encoding)
}

func (s *CityLibServer) middleware(w http.ResponseWriter, r *http.Request) error {

	s.setEncodingHeaders(w)

	return nil
}

func (s *CityLibServer) rootHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println(r.Method, r.URL.Path, r.Header.Get("Content-Type"))

	if handler, ok := s.handlers[r.URL.Path]; ok {
		if err := s.middleware(w, r); err != nil {
			return
		} else if err := handler(w, r); err != nil {
			http.Error(w, err.StatusText, err.StatusCode)
			return
		}
	}
}

func (s *CityLibServer) handleGetAllBooksRequest(w http.ResponseWriter, r *http.Request) *HttpErrorResponse {

	books, err := s.books.GetAll()

	if err != nil {
		s.logger.Println(err)
		return NewError(http.StatusInternalServerError)
	}

	if _, err := w.Write(books.AsJson()); err != nil {
		s.logger.Println(err)
		return NewError(http.StatusInternalServerError)
	}

	return nil
}

func (s *CityLibServer) handleInsertBookRequest(w http.ResponseWriter, r *http.Request) *HttpErrorResponse {
	var insertBookRequest requestmodels.InsertBookRequest

	bodyData, err := io.ReadAll(r.Body)

	if err != nil {
		s.logger.Println(err)
		return NewError(http.StatusBadRequest)

	}

	if err := json.Unmarshal(bodyData, &insertBookRequest); err != nil {
		s.logger.Println(err)
		return NewError(http.StatusBadRequest)
	}

	result, err := s.books.Insert(insertBookRequest)

	if err != nil {
		s.logger.Println(err)
		return NewError(http.StatusInternalServerError)
	}

	if _, err = w.Write(result.AsJson()); err != nil {
		s.logger.Println(err)
		return NewError(http.StatusInternalServerError)
	}

	return nil
}
