package repositories

import "rac_oblak_proj/models"

type Repository[T models.Querier] interface {
	Insert(t T) (T, error)
	GetAll() ([]T, error)
	GetById(id int64) (T, error)
	FilterBy(filter func(T) bool) ([]T, error)
	Remove(t T) error
}
