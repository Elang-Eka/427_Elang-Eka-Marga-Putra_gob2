package main

import (
	"fmt"
)

func main() {
	//Function
	menyapa("Elang", 22)

	//function bagian 2
	perkenalan("Elang", "jalan Sri Minulyo")
}

// funt menyapa
func menyapa(nama string, umur int8) { //parameter nama dan umur
	fmt.Println("==============================================================================")
	fmt.Printf("Halo semuanya, nama saya %s dan umur saya %d, salam kenal semua!!\n", nama, umur)
	fmt.Println("==============================================================================")
}

// func perkenalan
func perkenalan(nama string, alamat string) { //parameter nama dan alamat
	fmt.Printf("Halo semuanya, nama saya %s!!\n", nama)
	fmt.Printf("Tempat tinggal saya di %s, salam kenal semua!!\n", alamat)
	fmt.Println("==============================================================================")
}
