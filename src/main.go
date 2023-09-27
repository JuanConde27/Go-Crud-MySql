package main

import (
	"modcrudmysql.com/src/commons"
	"modcrudmysql.com/src/routes"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()
	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)
	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	log.Println("Servidor escuchando en http://localhost:3000")
	log.Println(server.ListenAndServe())
}