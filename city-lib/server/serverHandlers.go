package server

import (
	"net/http"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/errors/http_errors"
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
	} else if _, err := w.Write(books.AsJson()); err != nil {
		s.BaseServer.Logger.Println(err)
		return http_errors.NewError(http.StatusInternalServerError)
	}

	return nil
}

func (s *CityLibServer) handleInsertBookRequest(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	if r.Method != http.MethodPost {
		return http_errors.NewError(http.StatusMethodNotAllowed)
	}

	req, err := baseserver.ReadBody[requestmodels.InsertBookRequest](r)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest)
	}

	defer r.Body.Close()

	result, err := s.books.Insert(*req)

	if err != nil {
		s.BaseServer.Logger.Println(err)

		return http_errors.NewError(http.StatusInternalServerError)
	}

	if _, err = w.Write(result.AsJson()); err != nil {
		s.BaseServer.Logger.Println(err)

		return http_errors.NewError(http.StatusInternalServerError)
	}

	return nil
}
