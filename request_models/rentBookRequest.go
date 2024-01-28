package requestmodels

import (
	"encoding/json"
	"fmt"
)

type RentBookRequest struct {
	Username string `json:"username"`
	ISBN     string `json:"isb"`
}

func (r *RentBookRequest) String() string {
	return fmt.Sprintf("(%s, %s)", r.Username, r.ISBN)
}

func (r RentBookRequest) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}
