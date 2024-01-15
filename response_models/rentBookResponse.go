package responsemodels

import "encoding/json"

type RentBookResponse struct {
	Message string `json:"message"`
}

func (r *RentBookResponse) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}
