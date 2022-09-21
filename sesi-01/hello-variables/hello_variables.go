package main

import (
	"fmt"
)


func main() {
	// Belajar Variabel with data type
	var namaLengkap string
	var umur int

	namaLengkap = "Elang"
	umur = 22
	fmt.Println("Hello nama saya ", namaLengkap)
	fmt.Println("umur saya ", umur)
	fmt.Println("=====================================================================")

	// Belajar variabel without data type
	var namaAnda = "Elang"
	var umurAnda = 22
	fmt.Printf("%T , %T\n", namaAnda, umurAnda)
	fmt.Println("=====================================================================")

	// Belajar short declaration
	namaSaya := "Elang"
	umurSaya := 22
	fmt.Printf("%T , %T\n", namaSaya, umurSaya)
	fmt.Println("=====================================================================")

	// Multiple variabel declaration with data type
	var one, two, three string = "1", "2", "3"
	var angkaOne, angkaTwo, angkaThree int = 1, 2, 3
	fmt.Println("test data multiple => ", one, two, three)
	fmt.Println("test data multiple => ", angkaOne, angkaTwo, angkaThree)
	fmt.Println("=====================================================================")

	// Multiple variabel declaration without data type
	var one1, two2, three3 = "1", "2", "3"
	var angkaOne1, angkaTwo2, angkaThree3 = 1, 2, 3
	fmt.Println("test data multiple => ", one1, two2, three3)
	fmt.Println("test data multiple => ", angkaOne1, angkaTwo2, angkaThree3)
	fmt.Println("=====================================================================")

	// Multiple variabel declaration short
	dataOne, dataTwo, dataThree := "1", "2", "3"
	dataAngkaOne, dataAngkaTwo, dataAngkaThree := 1, 2, 3
	fmt.Println("test data multiple => ", dataOne, dataTwo, dataThree)
	fmt.Println("test data multiple => ", dataAngkaOne, dataAngkaTwo, dataAngkaThree)
	fmt.Println("=====================================================================")

	// Belajar underscore variabel & fmt printf usage
	var testVariabel string
	var oneName, twoName, alamatSaya, dataUmur = "Elang", "Eka", "Sidoarjo", 22
	_, _, _, _, _ = testVariabel, oneName, twoName, alamatSaya, dataUmur
	fmt.Printf("test underscore variabel => %T \n", oneName)
	fmt.Printf("halo nama saya %s, umur saya %d, saya berasal dari %s\n", oneName, dataUmur, alamatSaya)
	fmt.Println("=====================================================================")
}
