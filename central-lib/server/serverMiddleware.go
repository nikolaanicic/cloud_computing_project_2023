package server

import (
	"net/http"
	"rac_oblak_proj/errors/http_errors"
	"strings"
)

func (s *CentralLibServer) AllowedHost(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	host := strings.Split(r.RemoteAddr, ":")[0]

	allowedHosts := s.cfg.AllowedHosts

	for _, awh := range allowedHosts {
		if host == awh {
			return nil
		}
	}

	return http_errors.NewError(http.StatusBadRequest)
}

func (s *CentralLibServer) PostMethodAllowed(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	if r.Method != http.MethodPost {
		return http_errors.NewError(http.StatusMethodNotAllowed)
	}

	return nil
}

func (s *CentralLibServer) GetMethodAllowed(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	if r.Method != http.MethodGet {
		return http_errors.NewError(http.StatusMethodNotAllowed)
	}

	return nil
}
