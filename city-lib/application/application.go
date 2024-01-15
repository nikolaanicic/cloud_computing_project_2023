package application

import (
	"rac_oblak_proj/city-lib/config"
	"rac_oblak_proj/city-lib/data"
	"rac_oblak_proj/city-lib/repositories"

	"github.com/go-sql-driver/mysql"
)

type Application struct {
	config     *config.Config
	bookRepo   *repositories.BookRepo
	rentalRepo *repositories.RentalRepo
}

func New(config *config.Config) (*Application, error) {
	dataContext, err := data.NewDataContext(mysql.Config{
		User:      config.User,
		Passwd:    config.Password,
		Net:       "tcp",
		Addr:      config.DbHost,
		DBName:    config.DbName,
		ParseTime: true,
	})

	if err != nil {
		return nil, err
	}

	return &Application{
		config:     config,
		bookRepo:   repositories.NewBookRepo(dataContext),
		rentalRepo: repositories.NewRentalRepo(dataContext),
	}, nil
}

func (app *Application) Start() {

}
