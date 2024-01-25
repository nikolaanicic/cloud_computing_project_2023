package server

import (
	"net/http"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/errors/http_errors"
	requestmodels "rac_oblak_proj/request_models"
)

const (
	maxRent = 3
)

func (s *CentralLibServer) handleInsertUser(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	req, err := baseserver.ReadBody[requestmodels.InsertUserRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest)
	}

	defer r.Body.Close()

	result, err := s.userRepo.Insert(*req)

	if err != nil {
		s.BaseServer.Logger.Println(err)
		return http_errors.NewError(http.StatusConflict)
	}

	return baseserver.PackResponse(result, w, s.BaseServer.Logger)
}

func (s *CentralLibServer) handleUserLogin(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	req, err := baseserver.ReadBody[requestmodels.UserLoginRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest)
	}

	defer r.Body.Close()

	user, err := s.userRepo.GetByUsername(req.Username)

	if err != nil {
		return http_errors.NewError(http.StatusNotFound)
	} else if !s.userRepo.ValidatePassword(user.Password, user.Username, req.Password) {
		return http_errors.NewError(http.StatusUnauthorized)
	}

	s.BaseServer.Logger.Println("RESPONSE:", user.String())

	return baseserver.PackResponse(user, w, s.BaseServer.Logger)
}

func (c *CentralLibServer) handleRentBook(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	req, err := baseserver.ReadBody[requestmodels.RentBookRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest)
	}

	defer r.Body.Close()

	user, err := c.userRepo.GetByUsername(req.Username)

	if err != nil {
		return http_errors.NewError(http.StatusNotFound)
	} else if user.Rentals >= maxRent {
		return http_errors.NewError(http.StatusConflict)
	}

	user, err = c.userRepo.UpdateRentals(req.Username, 1)

	if err != nil {
		return http_errors.NewError(http.StatusInternalServerError)
	}

	return nil
}
