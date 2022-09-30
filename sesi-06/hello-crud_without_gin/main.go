package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// struct employee
type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

type Info struct {
	Message string `json:"message"`
}

// inisialisasi variabel employess
var employees = []Employee{
	{ID: 1, Name: "Elang", Age: 22, Division: "IT"},
	{ID: 2, Name: "Eka", Age: 22, Division: "Engineer"},
	{ID: 3, Name: "Putra", Age: 22, Division: "Staff"},
}

// penamaan port
var PORT = ":8080"

func main() {
	r := mux.NewRouter()

	// Route handlers
	r.HandleFunc("/getEmployees", getEmployees).Methods("GET")
	r.HandleFunc("/createEmployees", postEmployees).Methods("POST")
	r.HandleFunc("/getIdEmployees/{id}", getIdEmployee).Methods("GET")
	r.HandleFunc("/updateEmployees/{id}", updateEmployees).Methods("PUT")
	r.HandleFunc("/deleteEmployees/{id}", deleteEmployees).Methods("DELETE")

	fmt.Println("Application is listening on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}

// function get untuk mendapatkan value dari employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	/*
		Mengatur Content-Type menjadi application/json dalam method
		Set untuk mengirim response data dalam bentuk JSON
	*/
	w.Header().Set("Content-Type", "Application/json")

	// method get = request untuk mendapatkan data nilai atau value
	if r.Method == "GET" {
		// konversi data employess menjadi bentuk json
		info := Info{Message: "Get All Employess"}
		json.NewEncoder(w).Encode(info)
		json.NewEncoder(w).Encode(employees)
		return
	}

	// jika salah method maka akan muncul response error
	/*
		http.StatusBadRequest merupakan sebuah konstanta dari
		package http.StatusBadRequest yang merepresentasikan status
		400 yang merupakan code badRequest.
	*/
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

// function post/create data employees
func postEmployees(w http.ResponseWriter, r *http.Request) {
	/*
		Mengatur Content-Type menjadi application/json dalam method
		Set untuk mengirim response data dalam bentuk JSON
	*/
	w.Header().Set("Content-Type", "Application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid Age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)
		info := Info{Message: "Create Success"}
		json.NewEncoder(w).Encode(info)
		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

// Get By ID
func getIdEmployee(w http.ResponseWriter, r *http.Request) {

	/*
		Mengatur Content-Type menjadi application/json dalam method
		Set untuk mengirim response data dalam bentuk JSON
	*/
	w.Header().Set("Content-Type", "Application/json")

	// Get params
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid Id")
	}

	// method POST = request untuk mendapatkan data nilai atau value
	if r.Method == "GET" {
		jumlahEmployee := 0
		for i := 0; i < len(employees); i++ {
			jumlahEmployee = jumlahEmployee + 1
		}
		if employeeID > jumlahEmployee || employeeID == 0 {
			http.Error(w, "Id Not Found", http.StatusNotFound)
			return
		} else {
			for _, v := range employees {
				if employeeID == v.ID {
					// konversi data employess menjadi bentuk json
					info := Info{Message: "ID Found"}
					json.NewEncoder(w).Encode(info)
					json.NewEncoder(w).Encode(employees[employeeID-1])
					return
				}
			}
		}
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

// Function Update
func updateEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	name := r.FormValue("name")
	age := r.FormValue("age")
	division := r.FormValue("division")
	convertAge, err := strconv.Atoi(age)
	if err != nil {
		http.Error(w, "Invalid Age", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid Id")
	}
	if r.Method == "PUT" {
		jumlahEmployee := 0
		for i := 0; i < len(employees); i++ {
			jumlahEmployee = jumlahEmployee + 1
		}
		if employeeID > jumlahEmployee || employeeID == 0 {
			http.Error(w, "Id Not Found", http.StatusNotFound)
			return
		} else {
			for index, v := range employees {
				if employeeID == v.ID {
					employees = append(employees[:index], employees[index+1:]...)
					newEmployee := Employee{
						ID:       employeeID,
						Name:     name,
						Age:      convertAge,
						Division: division,
					}
					employees = append(employees, newEmployee)
					json.NewEncoder(w).Encode(employees[employeeID-1])
					info := Info{Message: "Update Success"}
					json.NewEncoder(w).Encode(info)
					return
				}
			}
		}
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

// Funtion Delete
func deleteEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid Id")
	}

	if r.Method == "DELETE" {
		jumlahEmployee := 0
		for i := 0; i < len(employees); i++ {
			jumlahEmployee = jumlahEmployee + 1
		}
		if employeeID > jumlahEmployee || employeeID == 0 {
			http.Error(w, "Id Not Found", http.StatusNotFound)
			return
		} else {
			for index, v := range employees {
				if employeeID == v.ID {
					json.NewEncoder(w).Encode(employees[employeeID-1])
					employees = append(employees[:index], employees[index+1:]...)
					info := Info{Message: "Delete Success"}
					json.NewEncoder(w).Encode(info)
					return
				}
			}
		}
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
