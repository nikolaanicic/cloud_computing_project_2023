package server

import (
	"fmt"
	"net/http"
	baseserver "rac_oblak_proj/base_server"
	"rac_oblak_proj/errors/http_errors"
	"rac_oblak_proj/models"
	requestmodels "rac_oblak_proj/request_models"
	"time"
)

func (s *CityLibServer) handleGetAllBooksRequest(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	books, err := s.books.GetAll()

	if err != nil {
		s.BaseServer.Logger.Println(err)
		return http_errors.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to get the books: %v", err))
	}

	return baseserver.PackResponse(books, w, s.BaseServer.Logger)
}

func (s *CityLibServer) handleInsertBookRequest(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	req, err := baseserver.ReadBody[requestmodels.InsertBookRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest, fmt.Sprintf("failed to insert a new book: %v", err))
	}

	defer r.Body.Close()

	result, err := s.books.Insert(*req)

	if err != nil {
		s.BaseServer.Logger.Println(err)
		return http_errors.NewError(http.StatusConflict, fmt.Sprintf("failed to insert a new book: %v", err))
	}

	return baseserver.PackResponse(result, w, s.BaseServer.Logger)
}

func (c *CityLibServer) handleUserLogin(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	req, err := baseserver.ReadBody[requestmodels.UserLoginRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest, fmt.Sprintf("failed to login the user: %v", err))
	}

	defer r.Body.Close()

	if c.sessionmgr.IsValid(req.Username) {
		token := getToken(r)
		c.sessionmgr.RefreshSession(token)

		return nil
	}

	response, err := baseserver.PostData(req, "http://"+c.config.CentralServerHost+"/users/login")

	if err != nil {
		return http_errors.NewError(http.StatusServiceUnavailable, fmt.Sprintf("failed to login the user: %v", err))
	}

	success := func() *http_errors.HttpErrorResponse {
		user, err := baseserver.ReadBody[models.User](response.Body)

		if err != nil {
			return http_errors.NewError(http.StatusBadRequest, fmt.Sprintf("failed to login the user: %v", err))
		}

		defer response.Body.Close()

		token := c.sessionmgr.AddSession(user)

		c.BaseServer.Logger.Println("RESPONSE:", "Token:", token.Value)

		return baseserver.PackResponse(token, w, c.BaseServer.Logger)
	}

	return baseserver.ParseResponse(response, success, c.BaseServer.GetReadHttpErrFunc(response.Body))
}

func (c *CityLibServer) handleReturnBook(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	req, err := baseserver.ReadBody[requestmodels.RentBookRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest, fmt.Sprintf("unable to read request: %v", err))
	}

	defer r.Body.Close()

	book, err := c.books.GetByISBN(req.ISBN)
	if err != nil {
		return http_errors.NewError(http.StatusNotFound, fmt.Sprintf("unable to return the book: %v", err))
	}

	token := getToken(r)

	user := c.sessionmgr.Get(token).User
	req.Username = user.Username

	response, err := baseserver.PostData(req, "http://"+c.config.CentralServerHost+"/books/return")

	if err != nil {
		return http_errors.NewError(http.StatusServiceUnavailable, fmt.Sprintf("unable to return the book: %v", err))
	}

	success := func() *http_errors.HttpErrorResponse {
		rental, err := c.rentals.GetByMemberAndBookId(user.ID, book.ID)

		if err != nil {
			return http_errors.NewError(http.StatusInternalServerError, fmt.Sprintf("unable to return the book: %v", err))
		} else if err := c.rentals.UpdateIsBookReturned(rental.ID, true); err != nil {
			c.BaseServer.Logger.Println(err)
			return http_errors.NewError(http.StatusBadRequest, fmt.Sprintf("unable to return the book: %v", err))
		}

		return nil
	}

	return baseserver.ParseResponse(response, success, c.BaseServer.GetReadHttpErrFunc(response.Body))
}

func (c *CityLibServer) handleRentBook(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	req, err := baseserver.ReadBody[requestmodels.RentBookRequest](r.Body)

	if err != nil {
		return http_errors.NewError(http.StatusBadRequest, fmt.Sprintf("failed to rent a book: %v", err))
	}

	defer r.Body.Close()

	book, err := c.books.GetByISBN(req.ISBN)
	if err != nil {
		return http_errors.NewError(http.StatusNotFound, fmt.Sprintf("failed to rent a book: %v", err))
	}

	token := getToken(r)

	user := c.sessionmgr.Get(token).User
	req.Username = user.Username

	response, err := baseserver.PostData(req, "http://"+c.config.CentralServerHost+"/books/rent")

	if err != nil {
		return http_errors.NewError(http.StatusServiceUnavailable, fmt.Sprintf("failed to rent a book: %v", err))
	}

	success := func() *http_errors.HttpErrorResponse {
		newRental := models.NewRental(user.ID, book.ID, time.Now(), false)

		rental, err := c.rentals.Insert(*newRental)

		if err != nil {
			return http_errors.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to rent a book: %v", err))
		}

		return baseserver.PackResponse(rental, w, c.BaseServer.Logger)
	}

	return baseserver.ParseResponse(response, success, c.BaseServer.GetReadHttpErrFunc(response.Body))
}
