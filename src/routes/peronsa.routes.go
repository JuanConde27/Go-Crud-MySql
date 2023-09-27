package routes

import "github.com/gorilla/mux"
import "modcrudmysql.com/src/controllers"

func SetPersonaRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/persona").Subrouter()
	subRouter.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRouter.HandleFunc("/find/{id}", controllers.GetById).Methods("GET")
	subRouter.HandleFunc("/register", controllers.Register).Methods("POST")
	subRouter.HandleFunc("/login", controllers.Login).Methods("POST")
	subRouter.HandleFunc("/update/{id}", controllers.Update).Methods("PUT")
	subRouter.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE")
}
