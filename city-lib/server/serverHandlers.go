package server

import (
	"encoding/json"
	"io"
	"net/http"
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

	var insertBookRequest requestmodels.InsertBookRequest

	bodyData, err := io.ReadAll(r.Body)

	if err != nil {
		s.BaseServer.Logger.Println(err)
		return http_errors.NewError(http.StatusBadRequest)

	}

	if err := json.Unmarshal(bodyData, &insertBookRequest); err != nil {
		s.BaseServer.Logger.Println(err)

		return http_errors.NewError(http.StatusBadRequest)
	}

	result, err := s.books.Insert(insertBookRequest)

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
