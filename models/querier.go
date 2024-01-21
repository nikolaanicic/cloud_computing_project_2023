package models

type Querier interface {
	FieldTypes() []string
	DataFields() []any
}

func SetFields(t interface{}, fields []any) {
	switch tp := t.(type) {
	case *Book:
		tp.SetFields(fields)
	case *Rental:
		tp.SetFields(fields)
	case *User:
		tp.SetFields(fields)
	}
}
