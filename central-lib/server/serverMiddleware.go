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
