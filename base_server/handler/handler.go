package baseserver

import (
	"net/http"
	"rac_oblak_proj/errors/http_errors"
)

type Handler func(http.ResponseWriter, *http.Request) *http_errors.HttpErrorResponse
