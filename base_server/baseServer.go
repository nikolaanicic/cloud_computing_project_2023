package baseserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rac_oblak_proj/base_server/pipeline"
	"rac_oblak_proj/data_context"
	"rac_oblak_proj/errors/http_errors"
	"rac_oblak_proj/mapper"
)

var encoding = "application/json"

type BaseServer struct {
	Logger    *log.Logger
	mux       *http.ServeMux
	host      string
	pipelines map[string]*pipeline.Pipeline
	ctx       *data_context.DataContext
}

func New(host string, logger *log.Logger, ctx *data_context.DataContext) *BaseServer {
	return &BaseServer{
		pipelines: make(map[string]*pipeline.Pipeline),
		host:      host,
		Logger:    logger,
		mux:       http.NewServeMux(),
		ctx:       ctx,
	}
}

func (s *BaseServer) setEncodingHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-Type", encoding)
}

func (s *BaseServer) isValidEncoding(r *http.Request, wanted string, method string) bool {
	return r.Header.Get("Content-Type") == wanted && r.Method == method
}

func (s *BaseServer) middleware(w http.ResponseWriter, r *http.Request) *http_errors.HttpErrorResponse {

	s.setEncodingHeaders(w)

	return nil
}

func ReadBody[T mapper.JsonModel](body io.ReadCloser) (*T, error) {
	var t T

	bodyData, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("unable to read request: %v", err)
	}

	if err := json.Unmarshal(bodyData, &t); err != nil {
		return nil, fmt.Errorf("unable to deserialize request to json: %v", err)
	}

	return &t, nil
}

func PackResponse[T mapper.JsonModel](response T, w http.ResponseWriter, logger *log.Logger) *http_errors.HttpErrorResponse {

	data := response.AsJson()

	if _, err := w.Write(data); err != nil {
		logger.Println("FAILURE", err)
		return http_errors.NewError(http.StatusInternalServerError, "failed to write the response")
	}

	return nil
}

func PostData[T mapper.JsonModel](data T, url string) (*http.Response, error) {
	return http.Post(url, encoding, bytes.NewBuffer(data.AsJson()))
}

func (b *BaseServer) GetReadHttpErrFunc(body io.ReadCloser) func() *http_errors.HttpErrorResponse {
	return func() *http_errors.HttpErrorResponse {
		data, err := ReadBody[http_errors.HttpErrorResponse](body)

		if err != nil {
			return http_errors.NewError(http.StatusInternalServerError, "failed to read internal response")
		}

		defer body.Close()

		return data
	}
}

func ParseResponse(response *http.Response, success func() *http_errors.HttpErrorResponse, failure func() *http_errors.HttpErrorResponse) *http_errors.HttpErrorResponse {
	if response.StatusCode == http.StatusOK {
		return success()
	} else {
		return failure()
	}
}

func (s *BaseServer) rootHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Println(r.Method, r.URL.Path, r.Header.Get("Content-Type"))

	writeHttpStatusError := func(err *http_errors.HttpErrorResponse) {
		s.Logger.Println(err)
		w.WriteHeader(err.StatusCode)
		w.Write(err.AsJson())
	}

	if err := s.middleware(w, r); err != nil {
		writeHttpStatusError(err)
		return
	}

	if pipeline, ok := s.pipelines[r.URL.Path]; ok {
		s.Logger.Println("found pipeline:", pipeline)

		if err := pipeline.Execute(w, r); err != nil {
			writeHttpStatusError(err)
		}
	} else {
		writeHttpStatusError(http_errors.NewError(http.StatusNotFound, fmt.Sprintf("pipeline %s doesn't exist", r.URL.Path)))
	}
}

func (bs *BaseServer) Serve() {
	bs.mux.HandleFunc("/", bs.rootHandler)
	bs.Logger.Println("listening on", bs.host)

	defer bs.ctx.Close()

	if err := http.ListenAndServe(bs.host, bs.mux); err != nil {
		bs.Logger.Fatal(err)

	}
}

func (bs *BaseServer) RegisterPipeline(pipeline *pipeline.Pipeline) error {

	if _, ok := bs.pipelines[pipeline.Path]; ok {
		return fmt.Errorf("handler for %s exists", pipeline.Path)
	}

	bs.pipelines[pipeline.Path] = pipeline

	bs.Logger.Println("registered pipeline:", pipeline)

	return nil
}

func (bs *BaseServer) RegisterMiddleware(path string, midFunc func(http.ResponseWriter, *http.Request) *http_errors.HttpErrorResponse) {
	if p, ok := bs.pipelines[path]; ok {
		p.RegisterMiddleware(midFunc)
	}

	bs.Logger.Println("registered middleware for:", path)
}
