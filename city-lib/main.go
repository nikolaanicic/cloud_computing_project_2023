package main

import (
	"fmt"
	"log"
	"os"
	"rac_oblak_proj/city-lib/config"
	"rac_oblak_proj/city-lib/data"
	"rac_oblak_proj/city-lib/repositories"
	requestmodels "rac_oblak_proj/request_models"

	"github.com/go-sql-driver/mysql"
)

func main() {

	filename := "app_config.json"

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	appCfg, err := config.LoadConfig(f)

	if err != nil {
		f.Close()
		log.Fatal()
	}

	f.Close()

	cfg := mysql.Config{
		User:      appCfg.User,
		Passwd:    appCfg.Password,
		Net:       "tcp",
		Addr:      appCfg.DbHost,
		DBName:    appCfg.DbName,
		ParseTime: true,
	}

	ctx, err := data.NewDataContext(cfg)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer ctx.Close()

	bookRepo := repositories.NewBookRepo(ctx)

	book := requestmodels.NewInsertBookRequest("thus spoke zarathustra", "nietzsche", "1232132131")

	res, err := bookRepo.Insert(&book)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
