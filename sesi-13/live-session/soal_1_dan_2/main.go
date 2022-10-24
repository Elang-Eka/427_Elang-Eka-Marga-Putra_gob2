package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type People struct {
	Nama         string `json:"nama"`
	Kode_Peserta int    `json:"kode"`
}

type Msg struct {
	Message string `json:"message"`
}

var People2 = []People{
	{Nama: "Elang Eka Marga Putra", Kode_Peserta: 427},
}

var PORT = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/helloworld", getPeople)
	r.HandleFunc("/elang/{id}", checkParam).Methods("GET")

	fmt.Println("Application is listening on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(People2)
	} else {
		http.Error(w, "Invalid method", http.StatusBadRequest)
	}
}

func checkParam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	peopleId, _ := strconv.Atoi(params["id"])
	if r.Method == "GET" {
		if peopleId == 1 {
			var messages = []Msg{
				{Message: "Hello World"},
			}
			json.NewEncoder(w).Encode(messages)
		} else if peopleId == 2 {
			var messages = []Msg{
				{Message: "Data Not Found"},
			}
			json.NewEncoder(w).Encode(messages)
		} else {
			var msg = ""
			json.NewEncoder(w).Encode(msg)
		}
	} else {
		http.Error(w, "Invalid method", http.StatusBadRequest)
	}
}
