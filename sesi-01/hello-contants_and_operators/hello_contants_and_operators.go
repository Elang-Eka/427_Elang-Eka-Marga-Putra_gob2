package main

import (
	"fmt"
)

func main() {

	// Belajar constant
	const fullName string = "Elang Eka Marga Putra"
	fmt.Printf("Nama lengkap saya adalah %s\n", fullName)
	fmt.Println("=====================================================================")

	// Belajar operators (aritmetics)
	var value1 = 2 + 4*5
	var value2 = (2 + 4) * 5
	fmt.Printf("Hasil value 1: %d\n", value1)
	fmt.Printf("Hasil value 2: %d\n", value2)
	fmt.Println("=====================================================================")

	// Belajar operators (relational)
	var kondisi1 bool = 2 < 3
	var kondisi2 bool = "elang" == "Elang"
	var kondisi3 bool = 5 != 3
	var kondisi4 bool = 10 <= 9
	fmt.Println("Hasil kondisi 1: ", kondisi1)
	fmt.Println("Hasil kondisi 2: ", kondisi2)
	fmt.Println("Hasil kondisi 3: ", kondisi3)
	fmt.Println("Hasil kondisi 4: ", kondisi4)
	fmt.Println("=====================================================================")

	// Belajar operators (logical)
	var benar = true
	var salah = false

	var salahDanBenar = salah && benar
	fmt.Printf("salah && benar \t(%t)\n", salahDanBenar)

	var salahAtauBenar = salah || benar
	fmt.Printf("salah && benar \t(%t)\n", salahAtauBenar)

	var reverse = !salah
	fmt.Printf("salah && benar \t(%t)\n", reverse)
	fmt.Println("=====================================================================")
}
