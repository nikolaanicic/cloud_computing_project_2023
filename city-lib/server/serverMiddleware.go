package server

import (
	"net/http"
	"rac_oblak_proj/city-lib/server/session"
	"rac_oblak_proj/errors/http_errors"
)

const (
	tokenHeader string = "X-Auth-Lib-Token"
)

func getToken(r *http.Request) string {
	return r.Header.Get(tokenHeader)
}

func (s *CityLibServer) Auth(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	token := getToken(r)

	if token == "" || !s.sessionmgr.Exists(token) {
		return http_errors.NewError(http.StatusUnauthorized)
	}

	return nil
}

func (s *CityLibServer) Session(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	token := getToken(r)

	if ss := s.sessionmgr.Get(token); ss != nil && !session.HasExpired(ss) {
		ss.Refresh()
		return nil
	}

	return http_errors.NewError(http.StatusUnauthorized)
}
