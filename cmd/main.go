package main

import (
	"flag"

	"github.com/maayarosama/Blogging_system/app"
	"github.com/rs/zerolog/log"
)

func main() {
	configPath := flag.String("configPath", "../config.json", "config file path")
	flag.Parse()
	a, err := app.NewApp(*configPath)
	err = a.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

}
