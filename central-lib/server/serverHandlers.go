package server

import (
	"fmt"
	"net/http"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/errors/http_errors"
	requestmodels "rac_oblak_proj/request_models"
)

const (
	maxRent = 3
)

func (s *CentralLibServer) handleUserSignUp(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	req, err := baseserver.ReadBody[requestmodels.UserSignUpRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest, "failed to read the request")
	}

	defer r.Body.Close()

	result, err := s.userRepo.Insert(*req)

	if err != nil {
		s.BaseServer.Logger.Println(err)
		return http_errors.NewError(http.StatusConflict, fmt.Sprintf("failed to sign up the user: %v", err))
	}

	return baseserver.PackResponse(result, w, s.BaseServer.Logger)
}

func (s *CentralLibServer) handleUserLogin(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	req, err := baseserver.ReadBody[requestmodels.UserLoginRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest, fmt.Sprintf("failed to log in the user: %v", err))
	}

	defer r.Body.Close()

	user, err := s.userRepo.GetByUsername(req.Username)

	if err != nil {
		return http_errors.NewError(http.StatusNotFound, fmt.Sprintf("failed to login the user: %v", err))
	} else if !s.userRepo.ValidatePassword(user.Password, user.Username, req.Password) {
		return http_errors.NewError(http.StatusUnauthorized, "invalid credentials")
	}

	s.BaseServer.Logger.Println("RESPONSE:", user.String())

	return baseserver.PackResponse(user, w, s.BaseServer.Logger)
}

func (c *CentralLibServer) handleRentBook(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	req, err := baseserver.ReadBody[requestmodels.RentBookRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest, fmt.Sprintf("failed to rent the book: %v", err))
	}

	defer r.Body.Close()

	user, err := c.userRepo.GetByUsername(req.Username)

	if err != nil {
		return http_errors.NewError(http.StatusNotFound, fmt.Sprintf("failed to rent the book: %v", err))
	} else if user.Rentals >= maxRent {
		return http_errors.NewError(http.StatusConflict, fmt.Sprintf("failed to rent the book: %v", err))
	}

	user, err = c.userRepo.UpdateRentals(req.Username, 1)

	if err != nil {
		return http_errors.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to rent the book: %v", err))
	}

	return nil
}

func (c *CentralLibServer) handleReturnBook(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	req, err := baseserver.ReadBody[requestmodels.RentBookRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest)
	}
	defer r.Body.Close()

	if user, err := c.userRepo.GetByUsername(req.Username); err != nil {
		return http_errors.NewError(http.StatusNotFound)
	} else if user.Rentals == 0 {
		return http_errors.NewError(http.StatusConflict)
	} else if _, err := c.userRepo.UpdateRentals(req.Username, -1); err != nil {
		return http_errors.NewError(http.StatusInternalServerError)
	}

	return nil
}
