package application

import (
	"log"
	"os"
	"rac_oblak_proj/interfaces"

	"rac_oblak_proj/config"
)

type Application struct {
	server interfaces.Server
	logger *log.Logger
}

func New(configFilename string, logger *log.Logger, server interfaces.Server) (*Application, error) {

	f, err := os.Open(configFilename)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	cfg := &config.Config{}

	if err := cfg.Load(f); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	srv, err := server.Configure(logger, cfg)

	if err != nil {
		return nil, err
	}

	return &Application{
		server: srv,
		logger: logger,
	}, nil
}

func (app *Application) Run() {
	app.server.Serve()
}
