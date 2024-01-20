package main

import (
	"log"
	"os"
	"rac_oblak_proj/city-lib/application"
)

func main() {

	filename := "app_config.json"
	logger := log.New(os.Stdout, "[CITY LIB] ", 0)

	app, err := application.New(filename, logger)

	if err != nil {
		logger.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
