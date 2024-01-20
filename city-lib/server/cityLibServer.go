package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"rac_oblak_proj/city-lib/repositories"
	"rac_oblak_proj/data_context"
	"rac_oblak_proj/interfaces"
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

func New() interfaces.Server {
	return &CityLibServer{
		mux:      http.NewServeMux(),
		handlers: make(map[string]func(http.ResponseWriter, *http.Request) *HttpErrorResponse),
	}
}

func (s *CityLibServer) Configure(logger *log.Logger, data *data_context.DataContext, host string) interfaces.Server {
	s.setLogger(logger)
	s.setBookRepo(repositories.NewBookRepo(data))
	s.setRentalsRepo(repositories.NewRentalRepo(data))
	s.setHost(host)

	return s

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

func (s *CityLibServer) setEncodingHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-Type", encoding)
}

func (s *CityLibServer) middleware(w http.ResponseWriter, r *http.Request) *HttpErrorResponse {

	s.setEncodingHeaders(w)

	if r.Method == http.MethodPost && r.Header.Get("Content-Type") != encoding {
		return NewError(http.StatusNotAcceptable)
	}

	return nil
}

func (s *CityLibServer) rootHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println(r.Method, r.URL.Path, r.Header.Get("Content-Type"))

	writeHttpStatusError := func(err *HttpErrorResponse) {
		http.Error(w, err.StatusText, err.StatusCode)
	}

	if handler, ok := s.handlers[r.URL.Path]; ok {
		if err := s.middleware(w, r); err != nil {
			writeHttpStatusError(err)
		} else if err := handler(w, r); err != nil {
			writeHttpStatusError(err)
		}
	}
}

func (s *CityLibServer) handleGetAllBooksRequest(w http.ResponseWriter, r *http.Request) *HttpErrorResponse {

	if r.Method != http.MethodGet {
		return NewError(http.StatusMethodNotAllowed)
	}

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

	if r.Method != http.MethodPost {
		return NewError(http.StatusMethodNotAllowed)
	}

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
