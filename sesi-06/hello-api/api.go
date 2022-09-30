package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// struct employee
type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
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
	// membuat route /employees
	http.HandleFunc("/getEmployees", getEmployees)
	http.HandleFunc("/createEmployees", postEmployees)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

// function get untuk mendapatkan value dari employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	// Bagian 1
	// /*
	// 	Mengatur Content-Type menjadi application/json dalam method
	// 	Set untuk mengirim response data dalam bentuk JSON
	// */

	// w.Header().Set("Content-Type", "Application/json")

	// // method get = request untuk mendapatkan data nilai atau value
	// if r.Method == "GET" {
	// 	// konversi data employess menjadi bentuk json
	// 	json.NewEncoder(w).Encode(employees)
	// 	return
	// }

	// // jika salah method maka akan muncul response error
	// /*
	// 	http.StatusBadRequest merupakan sebuah konstanta dari
	// 	package http.StatusBadRequest yang merepresentasikan status
	// 	400 yang merupakan code badRequest.
	// */
	// http.Error(w, "Invalid Method", http.StatusBadRequest)

	// Bagian 2
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)
		return
	}
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

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
