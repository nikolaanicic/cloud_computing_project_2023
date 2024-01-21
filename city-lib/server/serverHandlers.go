package server

import (
	"net/http"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/errors/http_errors"
	"rac_oblak_proj/models"
	requestmodels "rac_oblak_proj/request_models"
)

func (s *CityLibServer) handleGetAllBooksRequest(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	if r.Method != http.MethodGet {
		return http_errors.NewError(http.StatusMethodNotAllowed)
	}

	books, err := s.books.GetAll()

	if err != nil {
		s.BaseServer.Logger.Println(err)
		return http_errors.NewError(http.StatusInternalServerError)
	}

	return baseserver.PackResponse(books, w, s.BaseServer.Logger)
}

func (s *CityLibServer) handleInsertBookRequest(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	if r.Method != http.MethodPost {
		return http_errors.NewError(http.StatusMethodNotAllowed)
	}

	req, err := baseserver.ReadBody[requestmodels.InsertBookRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest)
	}

	defer r.Body.Close()

	result, err := s.books.Insert(*req)

	if err != nil {
		s.BaseServer.Logger.Println(err)
		return http_errors.NewError(http.StatusConflict)
	}

	return baseserver.PackResponse(result, w, s.BaseServer.Logger)
}

func (c *CityLibServer) handleUserLogin(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	if r.Method != http.MethodPost {
		return http_errors.NewError(http.StatusMethodNotAllowed)
	}

	req, err := baseserver.ReadBody[requestmodels.UserLoginRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest)
	}

	defer r.Body.Close()

	if _, ok := c.loggedInUsers[req.Username]; ok {
		return http_errors.NewError(http.StatusConflict)
	}

	response, err := baseserver.PostData(req, "http://"+c.config.CentralServerHost+"/users/login")

	// uraditi hendlanje odgovora po kodu
	// ako je 200 deserijalizovati korisnika
	// ako je neki drugi kod deserijalizovati HttpErrorResponse

	if err != nil {
		return http_errors.NewError(http.StatusServiceUnavailable)
	}

	user, err := baseserver.ReadBody[models.User](response.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest)
	}

	defer response.Body.Close()

	c.loggedInUsers[user.Username] = user

	return nil
}
