package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Rout() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", CreateEmployee).Methods("POST") // Corrected endpoint spelling
	log.Fatal(http.ListenAndServe(":8080", r))
	// Now the correct URL will be http://localhost:8080/employee
}
