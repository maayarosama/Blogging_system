package main

import (
	"flag"

	"github.com/rs/zerolog/log"

	"github.com/maayarosama/Blogging_system/routes"
)

func main() {
	configPath := flag.String("configPath", "../config.json", "config file path")
	flag.Parse()
	// Initialize new server
	server, err := routes.NewServer(*configPath)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	// start server
	err = server.Start()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

}
