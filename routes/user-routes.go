package routes

import (
	"github.com/gorilla/mux"
	"github.com/maayarosama/Blogging_system/controllers"
)

var RegisterUserStoreRoutes = func(router *mux.Router, c *controllers.Controller) {
	router.HandleFunc("/user", c.GetUsers).Methods("GET")
	router.HandleFunc("/user/signup", c.SignUp).Methods("POST")
	router.HandleFunc("/user/signin", c.SignIn).Methods("POST")

}
