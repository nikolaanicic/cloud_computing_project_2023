package responsemodels

import (
	"encoding/json"
	"fmt"
)

type InsertBookResponse struct {
	ID        int64  `json:"id"`
	BookTitle string `json:"name"`
	Writer    string `json:"writer"`
	ISBN      string `json:"isbn"`
}

func (r InsertBookResponse) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}

func (r *InsertBookResponse) String() string {
	return fmt.Sprintf("(%d, %s, %s, %s)", r.ID, r.BookTitle, r.Writer, r.ISBN)
}
