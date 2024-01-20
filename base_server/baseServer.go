package baseserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rac_oblak_proj/errors/http_errors"
	"rac_oblak_proj/mapper"
)

var encoding = "application/json"

type BaseServer struct {
	Logger   *log.Logger
	mux      *http.ServeMux
	host     string
	handlers map[string]func(http.ResponseWriter, *http.Request) *http_errors.HttpErrorResponse
}

func New(host string, logger *log.Logger) *BaseServer {
	return &BaseServer{
		handlers: make(map[string]func(http.ResponseWriter, *http.Request) *http_errors.HttpErrorResponse),
		host:     host,
		Logger:   logger,
		mux:      http.NewServeMux(),
	}
}

func (s *BaseServer) setEncodingHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-Type", encoding)
}

func (s *BaseServer) isValidEncoding(r *http.Request, wanted string, method string) bool {
	return r.Header.Get("Content-Type") == wanted && r.Method == method
}

func (s *BaseServer) middleware(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	s.setEncodingHeaders(w)

	if !s.isValidEncoding(r, encoding, http.MethodPost) && !s.isValidEncoding(r, "", http.MethodGet) {
		return http_errors.NewError(http.StatusNotAcceptable)
	}

	return nil
}

func ReadBody[T mapper.JsonModel](r *http.Request) (*T, error) {
	var t T

	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bodyData, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *BaseServer) rootHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Println(r.Method, r.URL.Path, r.Header.Get("Content-Type"))

	writeHttpStatusError := func(err *http_errors.HttpErrorResponse) {
		http.Error(w, err.StatusText, err.StatusCode)
	}

	if err := s.middleware(w, r); err != nil {
		writeHttpStatusError(err)
		return
	}

	if handler, ok := s.handlers[r.URL.Path]; ok {
		if err := handler(w, r); err != nil {
			writeHttpStatusError(err)
		}
	} else {
		writeHttpStatusError(http_errors.NewError(http.StatusNotFound))
	}
}

func (bs *BaseServer) Serve() {
	bs.mux.HandleFunc("/", bs.rootHandler)

	bs.Logger.Println("listening on", bs.host)
	if err := http.ListenAndServe(bs.host, bs.mux); err != nil {
		bs.Logger.Fatal(err)
	}
}

func (bs *BaseServer) RegisterHandler(path string, handler func(http.ResponseWriter, *http.Request) *http_errors.HttpErrorResponse) error {

	if _, ok := bs.handlers[path]; ok {
		return fmt.Errorf("handler for %s exists", path)
	}

	bs.handlers[path] = handler

	return nil
}
