package server

import "fmt"

var ErrBadReqeustMethod = func(method string) error {
	return fmt.Errorf("invalid method \"%s\"", method)
}

var ErrBadRequest = func(err error) error {
	return fmt.Errorf("bad request: %v", err)
}
