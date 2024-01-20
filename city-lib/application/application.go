package application

import (
	"log"
	"os"
	"rac_oblak_proj/city-lib/config"
	"rac_oblak_proj/city-lib/data"
	"rac_oblak_proj/city-lib/repositories"
	"rac_oblak_proj/city-lib/server"

	"github.com/go-sql-driver/mysql"
)

type Application struct {
	config *config.Config
	server *server.CityLibServer
	logger *log.Logger
	ctx    *data.DataContext
}

func New(configFilename string, logger *log.Logger) (*Application, error) {

	f, err := os.Open(configFilename)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	config, err := config.LoadConfig(f)

	if err != nil {
		return nil, err
	}

	ctx, err := data.NewDataContext(mysql.Config{
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

	srv := server.New().
		WithLogger(logger).
		WithBookRepo(repositories.NewBookRepo(ctx)).
		WithRentalsRepo(repositories.NewRentalRepo(ctx)).
		WithHost(config.ServerHost)

	return &Application{
		config: config,
		server: srv,
		logger: logger,
		ctx:    ctx,
	}, nil
}

func (app *Application) Run() error {
	app.server.Serve()
	return app.ctx.Close()
}
