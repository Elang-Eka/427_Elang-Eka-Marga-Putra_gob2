package main

import (
	"fmt"
)

func main() {

	//if else
	var currentyear = 2021
	fmt.Println("==================================================================")
	if age := currentyear - 1998; age < 17 {
		fmt.Printf("kamu belum boleh membuat kartu sim => %d\n", age)
	} else {
		fmt.Printf("Kamu sudah boleh membuat kartu\n")
	}
	fmt.Println("==================================================================")

	//Switch case
	var score = 8

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}
	fmt.Println("==================================================================")
	//Switch with relational operators
	var dataScore = 8

	switch {
	case dataScore == 8:
		fmt.Println("Perfect")
	case (dataScore < 8) && (dataScore >= 3):
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Study Harder")
			fmt.Println("You need to learn more")
		}
	}
	fmt.Println("==================================================================")
	//Switch (fallthrough keyword)
	var dataScore2 = 3

	switch {
	case dataScore2 == 8:
		fmt.Println("Perfect")
	case (dataScore2 < 8) && (dataScore2 >= 3):
		fmt.Println("Awesome")
		fallthrough
	case dataScore2 < 5:
		fmt.Println("It is ok")
	default:
		{
			fmt.Println("Study Harder")
			fmt.Println("You need to learn more")
		}
	}
	fmt.Println("==================================================================")

	//Nested Conditions
	var dataScore3 = 10

	if dataScore3 > 7 {
		switch dataScore3 {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice!!!")
		}
	} else {
		if dataScore3 == 5 {
			fmt.Println("Not Bad")
		} else if dataScore3 == 3 {
			fmt.Println("Keep Trying")
		} else {
			fmt.Println("You can do it")
			if dataScore3 == 0 {
				fmt.Println("Try Harder")
			}
		}
	}
	fmt.Println("==================================================================")
}
