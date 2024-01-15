package requestmodels

import (
	"encoding/json"
	"fmt"
)

type InsertBookRequest struct {
	BookTitle string `json:"name"`
	Writer    string `json:"writer"`
	ISBN      string `json:"isbn"`
}

func (r *InsertBookRequest) String() string {
	return fmt.Sprintf("(%s, %s, %s)", r.Writer, r.BookTitle, r.ISBN)
}

func (r InsertBookRequest) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}

func NewInsertBookRequest(title, writer, isbn string) InsertBookRequest {
	return InsertBookRequest{
		BookTitle: title,
		Writer:    writer,
		ISBN:      isbn,
	}
}
