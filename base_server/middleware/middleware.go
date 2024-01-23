package baseserver

import (
	"net/http"
	"rac_oblak_proj/errors/http_errors"
)

type Middleware func(http.ResponseWriter, *http.Request) *http_errors.HttpErrorResponse
