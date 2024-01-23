package pipeline

import (
	"net/http"
	h "rac_oblak_proj/base_server/handler"
	middleware "rac_oblak_proj/base_server/middleware"

	"rac_oblak_proj/errors/http_errors"
)

type Pipeline struct {
	Path       string
	handler    h.Handler
	middleware []middleware.Middleware
}

func New(path string, handler h.Handler) *Pipeline {
	return &Pipeline{
		Path:       path,
		handler:    handler,
		middleware: make([]middleware.Middleware, 0),
	}
}

func (p *Pipeline) RegisterMiddleware(fncs ...middleware.Middleware) {
	p.middleware = append(p.middleware, fncs...)
}

func (p *Pipeline) executeMiddleware(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	for _, mwf := range p.middleware {
		if err := mwf(w, r); err != nil {
			return err
		}
	}

	return nil
}

func (p *Pipeline) Execute(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {
	if err := p.executeMiddleware(w, r); err != nil {
		return nil
	}

	return p.handler(w, r)
}

func (p *Pipeline) String() string {
	return p.Path
}
