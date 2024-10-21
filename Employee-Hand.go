package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
    w.Header().Set("Content-Type", "application/json")
    
    var allDataGetEmpId Employee
    employeeID := mux.Vars(r)["eid"]

    id, err := strconv.Atoi(employeeID)
    if err != nil {
        log.Printf("Invalid employee ID: %v", err)
        http.Error(w, "Invalid employee ID", http.StatusBadRequest)
        return
    }

    log.Printf("Fetching employee with ID: %d", id)

    // Check if database connection works properly
    if err := Database.Where("id = ?", id).First(&allDataGetEmpId).Error; err != nil {
        log.Printf("Error fetching employee: %v", err)
        http.Error(w, "Employee not found", http.StatusNotFound)
        return
    }

    log.Printf("Employee data: %+v", allDataGetEmpId)

    if err := json.NewEncoder(w).Encode(allDataGetEmpId); err != nil {
        log.Printf("Error encoding JSON response: %v", err)
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}





func UpdateEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var allDataGetEmpId Employee
	
	 Database.First(mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&allDataGetEmpId)
	Database.Save(&allDataGetEmpId)
	json.NewEncoder(w).Encode(allDataGetEmpId)
}


