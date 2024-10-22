package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Rout() {
	r := mux.NewRouter()
	// r.HandleFunc("/getdata", GetEmployeeData).Methods("GET")
	// r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	// r.HandleFunc("/getdatas/{eid}", GetEmployeeDataID).Methods("GET")
	// r.HandleFunc("/getdatass/{eid}", UpdateEmployeeData).Methods("PUT")
	
	r.HandleFunc("/userdata",singup).Methods("POST")
	
	r.HandleFunc("/login",login).Methods("POST")
	r.HandleFunc("/decode", decodeToken).Methods("GET")


	// Corrected endpoint spelling
	log.Fatal(http.ListenAndServe(":8080", r))
	
	// Now the correct URL will be http://localhost:8080/employee
}
