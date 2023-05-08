package main

import (
	"github.com/rs/zerolog/log"

	"github.com/maayarosama/Blogging_system/routes"
)

func main() {
	// Initialize new server
	server, err := routes.NewServer("../config.json")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	// start server
	err = server.Start()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

}
