package data_errors

import "fmt"

var ErrEmptyQuery error = fmt.Errorf("query is empty")
var ErrInvalidQuery error = fmt.Errorf("invalid query")
