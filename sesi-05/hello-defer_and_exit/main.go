package main

import (
	"fmt"
	"os"
)

func main() {
	// defer #1
	defer fmt.Println("defer function starts to execute")
	fmt.Println("======================================")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go learning center")
	fmt.Println("======================================")

	// defer #2
	callDeferFunc()
	fmt.Println("Hai everyone")
	fmt.Println("======================================")
	// Exit
	defer fmt.Println("Invoke with defer")

	fmt.Println("Before exiting")
	fmt.Println("======================================")
	os.Exit(1)
	// ketika exit maka bagian ini tidak akan ditampilkan
	fmt.Println("Tidak tampil")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
