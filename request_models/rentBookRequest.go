package requestmodels

import (
	"encoding/json"
	"fmt"
)

type RentBookRequest struct {
	Username  string `json:"username"`
	BookTitle string `json:"name"`
	Writer    string `json:"writer"`
}

func (r *RentBookRequest) String() string {
	return fmt.Sprintf("(%s, %s)", r.BookTitle, r.Writer)
}

func (r RentBookRequest) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}
