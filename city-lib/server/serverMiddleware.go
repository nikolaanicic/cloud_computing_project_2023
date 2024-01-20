package server

import (
	"net/http"
	"rac_oblak_proj/errors/http_errors"
)

func (s *CityLibServer) setEncodingHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-Type", encoding)
}

func (s *CityLibServer) isValidEncoding(r *http.Request, wanted string, method string) bool {
	return r.Header.Get("Content-Type") == wanted && r.Method == method
}

func (s *CityLibServer) middleware(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	s.setEncodingHeaders(w)

	if !s.isValidEncoding(r, encoding, http.MethodPost) && !s.isValidEncoding(r, "", http.MethodGet) {
		return http_errors.NewError(http.StatusNotAcceptable)
	}

	return nil
}
