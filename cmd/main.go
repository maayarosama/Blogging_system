package main

import (
	"flag"

	"github.com/maayarosama/Blogging_system/app"
	"github.com/rs/zerolog/log"
)

func main() {
	configPath := flag.String("configPath", "../config.json", "config file path")
	flag.Parse()
	a, _ := app.NewApp(*configPath)
	db, _ := a.InitiateDB(a.Config.Database.Path)

	a.DB = db

	err := a.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

}
