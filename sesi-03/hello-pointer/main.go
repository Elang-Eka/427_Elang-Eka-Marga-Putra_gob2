package main

import (
	"fmt"
	"strings"
)

func main() {
	//Pointer
	var firstNumber *int
	var secondNumber *int
	_, _ = firstNumber, secondNumber
	//Pointer (memory address)
	var nomorSatu int = 4
	var nomorDua *int = &nomorSatu
	fmt.Println("==============================================")
	fmt.Println("Nilai nomorSatu : ", nomorSatu)
	fmt.Println("Alamat memori nomorSatu : ", &nomorSatu)

	fmt.Println("Nilai nomorDua : ", *nomorDua)
	fmt.Println("Alamat memori nomorDua : ", &nomorDua)
	fmt.Println("==============================================")

	// Pointer (Changing value through pointer)
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)

	fmt.Println("secondPerson (value) : ", *secondPerson)
	fmt.Println("secondPerson (memori address) : ", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 50), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)

	fmt.Println("secondPerson (value) : ", *secondPerson)
	fmt.Println("secondPerson (memori address) : ", secondPerson)
	fmt.Println("==============================================")
	// Pointer (Pointer as a parameter)
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
	fmt.Println("==============================================")
}

// func changeValue
func changeValue(number *int) {
	*number = 20
}
