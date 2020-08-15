package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vbanurag/go-rest-biolerplate/controllers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
