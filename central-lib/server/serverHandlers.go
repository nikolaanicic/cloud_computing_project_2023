package server

import (
	"net/http"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/errors/http_errors"
	requestmodels "rac_oblak_proj/request_models"
)

func (s *CentralLibServer) handleInsertUser(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	if r.Method != http.MethodPost {
		return http_errors.NewError(http.StatusMethodNotAllowed)
	}

	req, err := baseserver.ReadBody[requestmodels.InsertUserRequest](r)

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
