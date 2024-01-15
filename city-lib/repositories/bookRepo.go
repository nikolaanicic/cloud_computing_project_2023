package repositories

import (
	"rac_oblak_proj/city-lib/data"
	"rac_oblak_proj/mapper"
	"rac_oblak_proj/models"
	requestmodels "rac_oblak_proj/request_models"
	responsemodels "rac_oblak_proj/response_models"
)

type BookRepo struct {
	ctx *data.DataContext
}

func NewBookRepo(ctx *data.DataContext) *BookRepo {
	return &BookRepo{ctx}
}

func (r *BookRepo) Insert(book requestmodels.InsertBookRequest) (responsemodels.BookResponse, error) {

	newBook, err := mapper.Map[requestmodels.InsertBookRequest, models.Book](book)

	if err != nil {
		return responsemodels.BookResponse{}, err
	}

	query := "INSERT INTO books (name, writer, isbn) VALUES (?,?,?);"

	affected, err := data.ExecuteInsert[models.Book](r.ctx, query, newBook)

	if affected != 1 || err != nil {
		return responsemodels.BookResponse{}, err
	}

	return mapper.Map[models.Book, responsemodels.BookResponse](newBook)
}

func (r *BookRepo) GetAll() (responsemodels.GetAllBooksResponse, error) {
	result, err := r.getAll()

	if err != nil {
		return nil, err
	}

	return mapper.MapSlice[models.Book, responsemodels.BookResponse](result)
}

func (r *BookRepo) getAll() ([]models.Book, error) {
	return data.ExecuteQuery[models.Book](r.ctx, "SELECT * from books;")
}

func (r *BookRepo) FilterBy(filter func(b models.Book) bool) ([]models.Book, error) {

	all, err := r.getAll()

	if err != nil {
		return nil, err
	}

	result := make([]models.Book, 0)

	for _, b := range all {
		if filter(b) {
			result = append(result, b)
		}
	}

	return result, nil
}

func (r *BookRepo) GetById(id int64) (models.Book, error) {

	query := "SELECT * from books where id = ?;"

	result, err := data.ExecuteQuery[models.Book](r.ctx, query, []int64{id})

	if err != nil {
		return models.Book{}, err
	}

	return result[0], nil
}

func (r *BookRepo) Remove(b *models.Book) error {

	stmt := "REMOVE FROM books where id = ?;"

	affected, err := data.ExecuteStatement[models.Book](r.ctx, stmt, []int64{b.ID})

	if affected != 1 || err != nil {
		return err
	}

	return nil
}
