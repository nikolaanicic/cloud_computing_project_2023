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

	if c.sessionmgr.Exists(req.Username) {
		return http_errors.NewError(http.StatusConflict)
	}

	response, err := baseserver.PostData(req, "http://"+c.config.CentralServerHost+"/users/login")

	if err != nil {
		return http_errors.NewError(http.StatusServiceUnavailable)
	}

	success := func() *http_errors.HttpErrorResponse {
		user, err := baseserver.ReadBody[models.User](response.Body)

		if err != nil {
			return http_errors.NewError(http.StatusBadRequest)
		}

		defer response.Body.Close()

		c.sessionmgr.AddSession(user)

		c.BaseServer.Logger.Println("RESPONSE:", user.String())

		return nil
	}

	return baseserver.ParseResponse(response, success, c.BaseServer.GetReadHttpErrFunc(response.Body))
}
