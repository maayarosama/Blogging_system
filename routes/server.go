// Package routes for API endpoints
package routes

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/maayarosama/Blogging_system/config"
	"github.com/maayarosama/Blogging_system/controllers"
	"github.com/maayarosama/Blogging_system/models"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	host string
	port string
}

// Initialize the server and db
func NewServer(file string) (server *Server, err error) {
	f, err := config.ReadConfigFile(file)
	if err != nil {
		return
	}
	configuration, err := config.ParseConfigFile(f)
	if err != nil {
		return
	}

	db := models.NewDB()
	err = db.Connect(configuration.Database.File)
	if err != nil {
		return
	}
	err = db.Migrate()
	if err != nil {
		return
	}
	controller := controllers.NewController(db)
	r := mux.NewRouter()
	RegisterUserStoreRoutes(r, controller)
	// RegisterBlogStoreRoutes(r, controller)
	http.Handle("/", r)

	return &Server{port: configuration.Server.Port, host: configuration.Server.Host}, nil
}

// start the server and gracefull shutdown
func (s *Server) Start() (err error) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msgf("Server is listening on %s%s", s.host, s.port)

	srv := &http.Server{
		Addr: s.port,
	}

	go func() {
		// if err := srv.ListenAndServe(); err != nil {
		// 	log.Fatal().Err(err).Msg("HTTP server error")
		// }
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("HTTP server error")
		}
		log.Info().Msg("Stopped serving new connections")
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("HTTP shutdown error")
	}
	log.Info().Msg("Server is shutdown")

	return nil
}
