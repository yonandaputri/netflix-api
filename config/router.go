package config

import (
	"fmt"
	"log"
	"net/http"
	"netflixApp/tools"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func RunServer(router *mux.Router) {
	port := tools.ReadEnv("ROUTER_PORT", "3000")
	fmt.Println("Starting Web Server at port : " + port)
	err := http.ListenAndServe(": "+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
