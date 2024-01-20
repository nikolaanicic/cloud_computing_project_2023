package application

import (
	"log"
	"os"
	"rac_oblak_proj/city-lib/config"
	"rac_oblak_proj/city-lib/server"
	data "rac_oblak_proj/data_context"
	"rac_oblak_proj/interfaces"

	"github.com/go-sql-driver/mysql"
)

type Application struct {
	config interfaces.Config
	server interfaces.Server
	logger *log.Logger
	ctx    *data.DataContext
}

func New(configFilename string, logger *log.Logger) (*Application, error) {

	f, err := os.Open(configFilename)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	cfg := &config.Config{}

	if err := cfg.Load(f); err != nil {
		return nil, err
	}

	ctx, err := data.NewDataContext(mysql.Config{
		User:      cfg.User,
		Passwd:    cfg.Password,
		Net:       "tcp",
		Addr:      cfg.DbHost,
		DBName:    cfg.DbName,
		ParseTime: true,
	})

	if err != nil {
		return nil, err
	}

	srv := server.New().Configure(logger, ctx, cfg.ServerHost)

	return &Application{
		config: cfg,
		server: srv,
		logger: logger,
		ctx:    ctx,
	}, nil
}

func (app *Application) Run() error {
	app.server.Serve()
	return app.ctx.Close()
}
