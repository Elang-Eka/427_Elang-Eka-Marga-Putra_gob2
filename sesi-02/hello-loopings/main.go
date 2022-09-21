package main

import (
	"fmt"
)

func main() {
	//Loopings (first way of looping)
	for i := 0; i < 3; i++ {
		fmt.Println("Angka ", i)
	}
	fmt.Println("==============================================================")

	//Loopings (second way of looping)
	var datai = 0
	for datai < 3 {
		fmt.Println("Angka ", datai)
		datai++
	}
	fmt.Println("==============================================================")

	//Loopings (third way of looping)
	var data_i = 0
	for {
		fmt.Println("Angka ", data_i)

		data_i++
		if data_i == 3 {
			break
		}
	}
	fmt.Println("==============================================================")

	//Loopings (break and continue keywords)
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka ", i)
	}
	fmt.Println("==============================================================")

	//Loopings (Nested Looping)
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
	fmt.Println("==============================================================")

	//Loopings (Label)
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("perulangan ke - ", i+1)
		for j := i; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
