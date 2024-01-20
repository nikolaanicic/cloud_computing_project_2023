package main

import (
	"log"
	"os"
	"rac_oblak_proj/application"
	"rac_oblak_proj/city-lib/server"
)

func main() {

	filename := "app_config.json"
	logger := log.New(os.Stdout, "[CITY LIB] ", 0)

	srv := server.New()

	app, err := application.New(filename, logger, srv)

	if err != nil {
		logger.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
