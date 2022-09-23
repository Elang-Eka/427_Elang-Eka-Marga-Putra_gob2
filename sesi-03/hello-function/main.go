package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//Function
	menyapa("Elang Eka", 22)

	//function bagian 2
	perkenalan("Elang", "jalan Sri Minulyo")

	//function (return)
	var names = []string{"Elang", "Eka"}

	var printmsg = greet("Hai", names)

	fmt.Println(printmsg)
	fmt.Println("==============================================================================")

	// Function (Returning multiple values)
	var diameter float64 = 15

	var area, cicumference = calculate(diameter)

	fmt.Println("Area: ", area)
	fmt.Println("Circumference: ", cicumference)
	fmt.Println("==============================================================================")

	// Function (Predefined return value)
	var diameter2 float64 = 15

	var area2, cicumference2 = calculate2(diameter2)

	fmt.Println("Area: ", area2)
	fmt.Println("Circumference: ", cicumference2)
	fmt.Println("==============================================================================")

	//Function (Variadic function #1)
	studentLists := print("Elang", "Eka", "Marga", "Putra", "Artup")

	fmt.Printf("%s\n", &studentLists)
	fmt.Println("==============================================================================")

	// Function (Variadic function #2)
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberList...)

	fmt.Println("Result:", result)
	fmt.Println("==============================================================================")

	//Function (Variadic function #3)
	profile("Elang", "Mie", "Ayam Geprek", "Ikan Kerapu", "Sate Ponorogo")
	fmt.Println("==============================================================================")
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

// func greet
func greet(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// func calculate
func calculate(d float64) (float64, float64) {
	// luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	//keliling
	var circumference = math.Pi * d

	return area, circumference
}

// func calculate2
func calculate2(d float64) (area float64, circumference float64) {
	//luas
	area = math.Pi * math.Pow(d/2, 2)

	// keliling
	circumference = math.Pi * d

	return
}

// func print
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

// func sum
func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

// func profile
func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to set", mergeFavFoods)
}
