package models

import (
	"encoding/json"
	"time"
)

type Rental struct {
	ID             int64     `json:"id"`
	MemberID       int64     `json:"member_id"`
	BookID         int64     `json:"book_id"`
	RentalDate     time.Time `json:"rental_date"`
	IsBookReturned bool      `json:"is_book_returned" `
}

func (r Rental) FieldTypes() []string {
	return []string{"int64", "int64", "int64", "time.Time", "bool"}
}

func (r Rental) DataFields() []any {
	return []any{r.MemberID, r.BookID, r.RentalDate, r.IsBookReturned}
}

func (b Rental) String() string {
	val, _ := json.MarshalIndent(b, "", " ")

	return string(val)
}

func (b *Rental) SetFields(fields []any) {
	id, _ := fields[0].(*int64)
	b.ID = *id

	mib, _ := fields[1].(*int64)
	b.MemberID = *mib

	bib, _ := fields[2].(*int64)
	b.BookID = *bib

	rtime, _ := fields[3].(*time.Time)
	b.RentalDate = *rtime

	isReturned, _ := fields[4].(bool)
	b.IsBookReturned = isReturned
}

func NewRental(id, memberID, bookID int64, rentalDate time.Time, isBookReturned bool) *Rental {
	return &Rental{
		ID:             id,
		MemberID:       memberID,
		BookID:         bookID,
		RentalDate:     rentalDate,
		IsBookReturned: isBookReturned,
	}
}
