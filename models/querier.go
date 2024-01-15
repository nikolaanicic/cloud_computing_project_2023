package models

type Querier interface {
	FieldTypes() []string
	DataFields() []any
}

func SetFields(t interface{}, fields []any) {
	b, ok := (t).(*Book)

	if ok {
		b.SetFields(fields)
		return
	}

	r, ok := (t).(*Rental)

	if ok {
		r.SetFields(fields)
	}
}
