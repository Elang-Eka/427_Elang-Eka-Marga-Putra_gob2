package main

import (
	"fmt"
	"strings"
)

func main() {

	//array
	fmt.Println("==============================================")
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings1 = [3]string{"Elang", "Eka", "Putra"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings1)
	fmt.Println("==============================================")

	// Array(Modify Element Through Index)
	var fruits = [3]string{"Mangga", "Apel", "Pisang"}

	fruits[0] = "Manggo"
	fruits[1] = "Apple"
	fruits[2] = "banana"

	fmt.Printf("%#v\n", fruits)
	fmt.Println("==============================================")

	//Array(Loop through elements)
	var fruits2 = [3]string{"Mangga", "Apel", "Pisang"}

	//cara pertama
	for i, v := range fruits2 {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
	fmt.Println(strings.Repeat("#", 25))

	//cara kedua
	for i := 0; i < len(fruits2); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, fruits2[i])
	}
	fmt.Println("==============================================")

	// Array(Multidimensional Array)
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d", value)
		}
		fmt.Println()
	}
	fmt.Println("==============================================")

}
