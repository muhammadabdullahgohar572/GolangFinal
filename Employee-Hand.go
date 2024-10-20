package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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


	Database.Create(&empp)


	json.NewEncoder(w).Encode(empp)
}


func GetEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var allDataGetEmp []Employee
	
	 Database.Find(&allDataGetEmp)
	
	json.NewEncoder(w).Encode(allDataGetEmp)
}

func GetEmployeeDataID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var allDataGetEmpId Employee
	
	 Database.First(mux.Vars(r)["eid"])
	
	json.NewEncoder(w).Encode(allDataGetEmpId)
}


func UpdateEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var allDataGetEmpId Employee
	
	 Database.First(mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&allDataGetEmpId)
	Database.Save(&allDataGetEmpId)
	json.NewEncoder(w).Encode(allDataGetEmpId)
}


