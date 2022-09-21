package main

import (
	"fmt"
	"strings"
)

func main() {

	//slice
	var fruits = []string{"apple", "banana", "mango"}

	_ = fruits

	//slice (make function)
	var fruits2 = make([]string, 3)
	_ = fruits2
	fmt.Printf("%#v\n", fruits2)
	fmt.Println("====================================================================")
	//Slice (append function)
	var fruits3 = make([]string, 3)
	fruits3 = append(fruits3, "apple", "banana", "mango")
	fmt.Printf("%#v\n", fruits3)
	fmt.Println("====================================================================")
	//Slice (append function with ellipsis)

	var dataFruits1 = []string{"apple", "banana", "mango"}

	var dataFruits2 = []string{"durian", "pineapple", "starfruit"}

	dataFruits1 = append(dataFruits1, dataFruits2...)

	fmt.Printf("%#v\n", dataFruits1)
	fmt.Println("====================================================================")

	//Slice (copy function)
	var buah1 = []string{"apple", "banana", "mango"}

	var buah2 = []string{"durian", "pineapple", "starfruit"}

	nn := copy(buah1, buah2)

	fmt.Println("Buah1 => ", buah1)
	fmt.Println("Buah2 => ", buah2)
	fmt.Println("Copied elements => ", nn)
	fmt.Println("====================================================================")

	// Slice (slicing)
	var dataBuah1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var dataBuah2 = dataBuah1[0:4]
	fmt.Printf("%#v\n", dataBuah2)

	var dataBuah3 = dataBuah1[0:]
	fmt.Printf("%#v\n", dataBuah3)

	var dataBuah4 = dataBuah1[:3]
	fmt.Printf("%#v\n", dataBuah4)

	var dataBuah5 = dataBuah1[:]
	fmt.Printf("%#v\n", dataBuah5)
	fmt.Println("====================================================================")

	//Slice (Combining slicing and append)
	var fruitsData = []string{"apple", "banana", "mango", "durian", "pineapple"}
	fruitsData = append(fruitsData[:3], "rambutan")
	fmt.Printf("%#v\n", fruitsData)

	fmt.Println("====================================================================")

	//Slice (Backing array)
	var fruitsData1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruitsData2 = fruitsData1[2:4]

	fruitsData2[0] = "rambutan"

	fmt.Println("fruits1 => ", fruitsData1)
	fmt.Println("fruits2 => ", fruitsData2)
	fmt.Println("====================================================================")

	//Slice (Cap function)
	var dBuah1 = []string{"apple", "banana", "mango", "durian"}

	fmt.Println("fruits1 cap: ", cap(dBuah1))
	fmt.Println("fruits1 len: ", len(dBuah1))

	fmt.Println(strings.Repeat("#", 20))

	var dBuah2 = dBuah1[0:3]

	fmt.Println("fruits1 cap: ", cap(dBuah2))
	fmt.Println("fruits1 len: ", len(dBuah2))

	fmt.Println(strings.Repeat("#", 20))

	var dBuah3 = dBuah1[1:]

	fmt.Println("fruits1 cap: ", cap(dBuah3))
	fmt.Println("fruits1 len: ", len(dBuah3))
	fmt.Println("====================================================================")

	//Slice (Creating a new backing array)
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars: ", cars)
	fmt.Println("newcars: ", newCars)
	fmt.Println("====================================================================")
}
