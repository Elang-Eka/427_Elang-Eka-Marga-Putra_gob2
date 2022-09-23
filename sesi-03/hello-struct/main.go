package main

import (
	"fmt"
	"strings"
)

// Struct 1
type Employee struct {
	name     string
	age      int
	division string
}

// Struct 2
type Person struct {
	name string
	age  int
}

// Struct 3
type dataE struct {
	division string
	person   Person
}

func main() {
	// Struct (Giving value to struct)
	var employee Employee

	employee.name = "Elang"

	employee.age = 22

	employee.division = "Software Developer"

	fmt.Println("=============================================")
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
	fmt.Println("=============================================")

	// Struct (Initializing struct)
	var employee1 = Employee{}
	employee1.name = "Elang"

	employee1.age = 22

	employee1.division = "Software Developer"

	var employee2 = Employee{name: "Eka", age: 22, division: "Backend Developer"}

	fmt.Printf("Employee1 %+v\n", employee1)
	fmt.Printf("Employee2 %+v\n", employee2)
	fmt.Println("=============================================")

	// Struct (Pointer to a struct)
	var dataEmployee1 = Employee{name: "Elang", age: 22, division: "Backend Developer"}

	var dataEmployee2 *Employee = &dataEmployee1

	fmt.Println("Employee1 name:", dataEmployee1.name)
	fmt.Println("Employee2 name:", dataEmployee2.name)

	fmt.Println(strings.Repeat("#", 50))

	dataEmployee2.name = "Eka"

	fmt.Println("Employee1 name:", dataEmployee1.name)
	fmt.Println("Employee2 name:", dataEmployee2.name)
	fmt.Println("=============================================")

	// Struct (Embedded struct)
	var employ = dataE{}

	employ.person.name = "Elang"
	employ.person.age = 22
	employ.division = "Backend Developer"

	fmt.Printf("%+v\n", employ)
	fmt.Println("=============================================")

	// Struct (Anonymous struct)
	// Tanpa pengisian field
	var employ1 = struct {
		person   Person
		division string
	}{}

	employ1.person = Person{name: "Elang", age: 22}
	employ1.division = "Backend Developer"

	// dengan pengisian field

	var employ2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Eka", age: 22},
		division: "Fullstack",
	}

	fmt.Printf("Employee1: %+v\n", employ1)
	fmt.Printf("Employee2: %+v\n", employ2)
	fmt.Println("=============================================")

	// Struct (Slice of struct)
	var people = []Person{
		{name: "Elang", age: 22},
		{name: "Eka", age: 23},
		{name: "Marga", age: 24},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("=============================================")

	// Struct (Slice of anonymous struct)
	var emp = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Elang", age: 22}, division: "Backend Developer"},
		{person: Person{name: "Eka", age: 23}, division: "Frontend Developer"},
		{person: Person{name: "Marga", age: 24}, division: "Fullstack Developer"},
	}

	for _, v := range emp {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("=============================================")
}
