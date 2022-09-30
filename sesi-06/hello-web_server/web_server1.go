package main

import (
	"fmt"
	"net/http"
)

// alamat port
var PORT = ":8080"

func main() {
	// handleFunc/route
	http.HandleFunc("/", greet)

	// menjalankan serve
	http.ListenAndServe(PORT, nil)
}

// http.ResponseWriter = untuk mengirim response balik kepada pengirim req
// *http.Request = struct untuk mendapatkan data request yang di post oleh header, url parameter dan lainnya
func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Halo ini adalah web server golang"
	fmt.Fprint(w, msg)
}
