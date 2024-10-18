package main

import "gorm.io/gorm"

type Employeee struct {
    gorm.Model
    EmployeeName  string  `json:"EmpName"`
    EmployeeSalary float64 `json:"EmpSalary"`
    Email         string  `json:"Email"`
}
