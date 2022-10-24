package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type People struct {
	Nama         string `json:"nama"`
	Kode_Peserta int    `json:"kode"`
}

var dataDiri = []People{
	{Nama: "Elang Eka Marga Putra", Kode_Peserta: 427},
}

func main() {
	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(greet)

	mux.Handle("/", middleware1(middleware2(endpoint)))

	fmt.Println("Listening port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}

func greet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dataDiri)
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := "Hello Elang"
		json.NewEncoder(w).Encode(message)
		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := "Not Found"
		json.NewEncoder(w).Encode(message)
		next.ServeHTTP(w, r)
	})
}
