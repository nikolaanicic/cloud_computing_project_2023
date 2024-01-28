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
		return http_errors.NewError(http.StatusUnauthorized, "Wrong password or username")
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
		return http_errors.NewError(http.StatusUnauthorized, "please log in again")
	}
}

func (s *CityLibServer) PostMethodAllowed(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	if r.Method != http.MethodPost {
		return http_errors.NewError(http.StatusMethodNotAllowed, "method not allowed")
	}

	return nil
}

func (s *CityLibServer) GetMethodAllowed(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	if r.Method != http.MethodGet {
		return http_errors.NewError(http.StatusMethodNotAllowed, "method not allowed")
	}

	return nil
}
