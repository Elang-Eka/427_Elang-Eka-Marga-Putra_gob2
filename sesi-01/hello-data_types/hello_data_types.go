package main

import (
	"fmt"
)


func main() {
	// Belajar integer with declare data type
	var satu uint8 = 32
	var dua int8 = 45
	fmt.Printf("Tipe data satu => %T \n", satu)
	fmt.Printf("Tipe data dua  => %T \n", dua)
	fmt.Println("=====================================================================")

	// Belajar integer without declare data type
	var dataSatu = 32
	var dataDua = 45
	fmt.Printf("Tipe data satu => %T \n", dataSatu)
	fmt.Printf("Tipe data dua  => %T \n", dataDua)
	fmt.Println("=====================================================================")

	// Belajar float
	var dataDecimal = 32.5

	fmt.Printf("Angka decimal : %f \n", dataDecimal)
	fmt.Printf("Angka decimal : %.3f \n", dataDecimal)
	fmt.Println("=====================================================================")

	// Belajar boolean( true / false)
	var kondisi bool = true
	fmt.Printf("is it permitted? %T\n", kondisi)
	fmt.Println("=====================================================================")

	// Belajar string
	var message string = "halo"
	fmt.Println(message)
	fmt.Println("=====================================================================")

	// Belajar nil and zero value
	var data1 string
	var data2 bool
	var data3 int
	var data4 float64

	var data5 []int
	var data6 *int
	var data7 map[string]string
	var data8 func()
	var data9 interface{}
	// ZERO VALUE
	fmt.Println("string types: ", data1)
	fmt.Println("boolean types: ", data2)
	fmt.Println("integer types: ", data3)
	fmt.Println("float64 types: ", data4)

	//NIL
	fmt.Println("slice types: ", data5)
	fmt.Println("pointer types: ", data6)
	fmt.Println("map types: ", data7)
	fmt.Print("function types: ", data8)
	fmt.Println("interface types: ", data9)
	fmt.Println("=====================================================================")

	// Belajar string backticks
	var dataMessage = `
		Halo Perkenalkan nama saya
		"Elang Eka Marga Putra". 
		Saya merupakan lulusan dari
		Informatika UPN "Veteran" Jawa Timur.
		`
	fmt.Println(dataMessage)
	fmt.Println("=====================================================================")

}
