package responsemodels

import (
	"encoding/json"
	"fmt"
)

type BookResponse struct {
	BookTitle string `json:"name"`
	Writer    string `json:"writer"`
	ISBN      string `json:"isbn"`
}

func (r BookResponse) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}

func (r *BookResponse) String() string {
	return fmt.Sprintf("(%s, %s, %s)", r.BookTitle, r.Writer, r.ISBN)
}
