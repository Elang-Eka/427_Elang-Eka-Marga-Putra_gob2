package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//Looping Over String (byte-by-byte) bagian 1
	city := "Surabaya"

	for i := 0; i < len(city); i++ {
		fmt.Printf("Index: %d, byte: %d\n", i, city[i])
	}
	fmt.Println("======================================================")

	//Looping Over String (byte-by-byte) bagian 2
	var city2 []byte = []byte{74, 97, 107, 97, 114, 116, 97}

	fmt.Println(string(city2))
	fmt.Println("======================================================")

	//Looping Over String (rune-by-rune) bagian 1
	str1 := "makan"

	str2 := "mânca"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))
	fmt.Println("======================================================")

	//Looping Over String (rune-by-rune) bagian 2
	datastr1 := "makan"

	datastr2 := "mânca"

	fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(datastr1))
	fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(datastr2))
	fmt.Println("======================================================")

	//Looping Over String (rune-by-rune) bagian 3
	datastr := "mânca"

	for i, s := range datastr {
		fmt.Printf("Index => %d, string => %s\n", i, string(s))
	}
	fmt.Println("======================================================")
}
