package models

import (
	"encoding/json"
)

type Book struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Writer string `json:"writer"`
	ISBN   string `json:"isbn"`
}

func (b Book) AsJson() []byte {
	data, _ := json.Marshal(b)

	return data
}

func (b *Book) SetFields(fields []any) {
	id, _ := fields[0].(*int64)
	b.ID = *id

	name, _ := fields[1].(*string)
	b.Name = *name

	writer, _ := fields[2].(*string)
	b.Writer = *writer

	isbn, _ := fields[3].(*string)
	b.ISBN = *isbn

}

func (b Book) FieldTypes() []string {
	return []string{"int64", "string", "string", "string"}
}

func (b Book) DataFields() []any {
	return []any{b.Name, b.Writer, b.ISBN}
}

func (b Book) String() string {
	val, _ := json.MarshalIndent(b, "", " ")

	return string(val)
}

func NewBook(id int64, name, writer, isbn string) *Book {
	return &Book{
		ID:     id,
		Name:   name,
		Writer: writer,
		ISBN:   isbn,
	}
}
