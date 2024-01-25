package server

import (
	"net/http"
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

	if token == "" || s.sessionmgr.HasExpired(token) {
		return http_errors.NewError(http.StatusUnauthorized)
	}

	return nil
}

func (s *CityLibServer) Session(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	token := getToken(r)

	if s.sessionmgr.IsValid(token) {
		s.sessionmgr.RefreshSession(token)
		return nil
	} else {
		s.sessionmgr.RemoveSession(token)
		return http_errors.NewError(http.StatusUnauthorized)
	}
}
