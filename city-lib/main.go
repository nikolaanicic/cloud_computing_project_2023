package main

import (
	"log"
	"os"
	"rac_oblak_proj/city-lib/config"
	"rac_oblak_proj/city-lib/data"
	"rac_oblak_proj/city-lib/repositories"
	"rac_oblak_proj/city-lib/server"

	"github.com/go-sql-driver/mysql"
)

func main() {

	logger := log.New(os.Stdout, "[CITY LIB] ", 0)
	filename := "app_config.json"

	f, err := os.Open(filename)
	if err != nil {
		logger.Fatal(err)
	}

	appCfg, err := config.LoadConfig(f)

	if err != nil {
		f.Close()
		logger.Fatal(err)
	}

	f.Close()

	cftg := mysql.Config{
		User:      appCfg.User,
		Passwd:    appCfg.Password,
		Net:       "tcp",
		Addr:      appCfg.DbHost,
		DBName:    appCfg.DbName,
		ParseTime: true,
	}

	ctx, err := data.NewDataContext(cftg)

	if err != nil {
		logger.Fatal(err)
	}

	defer ctx.Close()

	srv := server.New().
		WithLogger(logger).
		WithBookRepo(repositories.NewBookRepo(ctx)).
		WithRentalsRepo(repositories.NewRentalRepo(ctx)).
		WithHost(appCfg.ServerHost)

	srv.Serve()
}
