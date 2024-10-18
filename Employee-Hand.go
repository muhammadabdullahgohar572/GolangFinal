package main

import (
	"encoding/json"
	"net/http"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	EmployeeName   string  `json:"EmpName"`
	EmployeeSalary float64 `json:"EmpSalary"`
	Email          string  `json:"Email"`
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var empp Employee
	json.NewDecoder(r.Body).Decode(&empp)

	// Save to the database
	Database.Create(&empp)

	// Return the created employee
	json.NewEncoder(w).Encode(&empp)
}

// You can uncomment and implement these when needed
// func GetEmployee(w http.ResponseWriter, r *http.Request) {}
// func GetEmployeeID(w http.ResponseWriter, r *http.Request) {}
