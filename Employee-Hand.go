package main

import (
	
	"encoding/json"
	"log"
	"net/http"

	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)
var jwtKey = []byte("abdullah")

// Define the Claims struct with the fields you need in the token
type Claims struct {
	Email string `json:"Email"`
	jwt.StandardClaims
}
type Employee struct {
	gorm.Model
	EmployeeName   string  `json:"EmpName"`
	EmployeeSalary float64 `json:"EmpSalary"`
	Email          string  `json:"Email"`
}

type CreateUserData struct {
	gorm.Model
	UserName string `json:"UserName"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Age      int    `json:"Age"`
	Gender   string `json:"Gender"`
	jwt.StandardClaims
}


func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func HashPassword(password string)(string,error)  {
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes),err
}

func singup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var userdatac CreateUserData
	json.NewDecoder(r.Body).Decode(&userdatac)
	HashPassword,err :=HashPassword(userdatac.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	userdatac.Password = HashPassword
    Database.Create(&userdatac)
	// userdatac.Password = ""
	json.NewEncoder(w).Encode(&userdatac)

}


func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userdatac CreateUserData
	var dbUser CreateUserData

	// Decode incoming request for email and password
	json.NewDecoder(r.Body).Decode(&userdatac)

	// Find user by email in the database
	if err := Database.Where("email = ?", userdatac.Email).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Check password hash
	if !CheckPasswordHash(userdatac.Password, dbUser.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Create JWT token with all user data
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
	claims := &CreateUserData{
		UserName: dbUser.UserName,
		Email:    dbUser.Email,
		Password: dbUser.Password, // Include only if necessary (not recommended)
		Age:      dbUser.Age,
		Gender:   dbUser.Gender,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Return the token
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}







func decodeToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the token from the URL query parameter
	tokenStr := r.URL.Query().Get("token")

	// Parse the token and validate its signature
	claims := &CreateUserData{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil // Replace jwtKey with your secret key
	})

	// Check if there was an error in decoding or if the token is invalid
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Return user data based on the token
	userData := map[string]interface{}{
		"UserName": claims.UserName,
		"Email":    claims.Email,
		"Password": claims.Password, // Include only if necessary
		"Age":      claims.Age,
		"Gender":   claims.Gender,
	}

	// Return the extracted user data as a JSON response
	json.NewEncoder(w).Encode(userData)
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
