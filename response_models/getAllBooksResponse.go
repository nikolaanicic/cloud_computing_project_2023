package responsemodels

import "encoding/json"

type GetAllBooksResponse []BookResponse

func (r GetAllBooksResponse) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}
