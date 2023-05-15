package app

// Ignore this for now

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maayarosama/Blogging_system/internal"
	"github.com/maayarosama/Blogging_system/models"

	"github.com/rs/zerolog/log"
)

type App struct {
	server *Server
	router *mux.Router
	config internal.Configuration
	db     models.DB

	// logger Logger
}

// Initiate and connect the db , register the routes and initiate the server from config.json file
func NewApp(path string) (a *App, err error) {

	f, err := internal.ReadConfigFile(path)
	if err != nil {
		return
	}
	config, err := internal.ParseConfigFile(f)
	if err != nil {
		return
	}

	db := models.NewDB()
	err = db.Connect(config.Database.Path)
	if err != nil {
		return
	}
	err = db.Migrate()
	if err != nil {
		return
	}

	//Initiate the app instance with data read from config.json
	a = &App{
		server: &Server{},
		router: mux.NewRouter(),
		config: config,
		db:     db,
	}

	a.RegisterHandlers()
	http.Handle("/", a.router)

	server, err := NewServer(config.Server.Port, config.Server.Host)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	a.server = server

	return a, nil
}

// Starts the server initiated in the app
func (a *App) ListenAndServe() error {
	return a.server.Start()
}

// Registering all routes
func (a *App) RegisterHandlers() {
	a.router.HandleFunc("/user", a.GetUsers).Methods("GET")
	a.router.HandleFunc("/user/signup", a.SignUp).Methods("POST")
	a.router.HandleFunc("/user/signin", a.SignIn).Methods("POST")
	a.router.HandleFunc("/user/verifyemail", a.VerifyEmail).Methods("POST")

	// a.router.HandleFunc("/blog/", a.CreateBlog).Methods("POST")
	// a.router.HandleFunc("/blog/", a.GetBlogs).Methods("GET")
	// a.router.HandleFunc("/blog/{BlogId}", a.GetBlogById).Methods("GET")
	// a.router.HandleFunc("/blog/{BlogId}", a.UpdateBlog).Methods("PUT")
	// a.router.HandleFunc("/blog/{BlogId}", a.DeleteBlog).Methods("DELETE")
}
