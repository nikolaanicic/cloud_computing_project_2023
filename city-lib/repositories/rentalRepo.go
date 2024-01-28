package repositories

import (
	"rac_oblak_proj/data_context"
	data "rac_oblak_proj/data_context"
	"rac_oblak_proj/models"
)

type RentalRepo struct {
	ctx *data.DataContext
}

func NewRentalRepo(ctx *data.DataContext) *RentalRepo {
	return &RentalRepo{ctx}
}

func (r *RentalRepo) Insert(rental models.Rental) (models.Rental, error) {

	query := "INSERT INTO rentals (memberid, bookid, rentaldate, isbookreturned) VALUES (?,?,?,?);"

	affected, err := data.ExecuteInsert[models.Rental](r.ctx, query, rental)

	if affected != 1 || err != nil {
		return models.Rental{}, err
	}

	return rental, nil
}

func (r *RentalRepo) GetAll() ([]models.Rental, error) {
	return data.ExecuteQuery[models.Rental](r.ctx, "SELECT * from rentals;")
}

func (r *RentalRepo) FilterBy(filter func(b models.Rental) bool) ([]models.Rental, error) {

	all, err := r.GetAll()

	if err != nil {
		return nil, err
	}

	result := make([]models.Rental, 0)

	for _, b := range all {
		if filter(b) {
			result = append(result, b)
		}
	}

	return result, nil
}

func (r *RentalRepo) GetById(id int64) (models.Rental, error) {
	query := "SELECT * from rentals where id = ?;"

	result, err := data.ExecuteQuery[models.Rental](r.ctx, query, []int64{id})

	if err != nil {
		return models.Rental{}, err
	}

	return result[0], nil
}

func (r *RentalRepo) GetByMemberAndBookId(memberId, bookId int64) (models.Rental, error) {
	query := "SELECT * from rentals where memberid = ? and bookid ? ORDER BY rentaldate desc"

	result, err := data.ExecuteQuery[models.Rental](r.ctx, query, memberId, bookId)

	if err != nil {
		return models.Rental{}, err
	}

	return result[0], nil
}

func (r *RentalRepo) IsBookAvailable(bookID int64) (bool, error) {
	retval := false

	query := "SELECT * from rentals where bookid = ? ORDER BY rentaldate desc"

	result, err := data.ExecuteQuery[models.Rental](r.ctx, query, bookID)

	if err != nil {
		return retval, err
	}

	return result[0].IsBookReturned, nil
}

func (r *RentalRepo) UpdateIsBookReturned(rentalID int64, value bool) error {

	stmt := "UPDATE rentals SET isbookreturned = ? where id = ?"

	_, err := data_context.ExecuteStatement(r.ctx, stmt, value, rentalID)

	if err != nil {
		return err
	}

	return nil
}

func (r *RentalRepo) GetByMemberId(memberId int64) ([]models.Rental, error) {

	return data.ExecuteQuery[models.Rental](
		r.ctx,
		"SELECT * from books WHERE memberid = ?;",
		[]int64{memberId})
}

func (r *RentalRepo) Remove(b *models.Rental) error {

	stmt := "REMOVE FROM rentals where id = ?;"

	affected, err := data.ExecuteStatement(r.ctx, stmt, []int64{b.ID})

	if affected != 1 || err != nil {
		return err
	}

	return nil
}
