package main

import (
	"fmt"
)

func main() {
	// defer #2
	callDeferFunc()
	fmt.Println("Hai everyone")
	fmt.Println("======================================")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute #2")
}
