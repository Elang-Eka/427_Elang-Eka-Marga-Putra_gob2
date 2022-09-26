package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func main() {
	// Identifying Data Type & Value
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("===================================================")
	fmt.Println("Tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variabel :", reflectValue.Int())
	}
	fmt.Println("===================================================")
	// Accessing Value Using Interface Method
	fmt.Println("Tipe variable :", reflectValue.Type())
	fmt.Println("Tipe variable :", reflectValue.Interface())

	nilai := reflectValue.Interface().(int)

	_ = nilai

	// Identifying Method Information
	fmt.Println("===================================================")
	var s1 = &student{Name: "Elang Eka", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue1 = reflect.ValueOf(s1)
	var method = reflectValue1.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Eka"),
	})

	fmt.Println("nama :", s1.Name)
	fmt.Println("===================================================")
}

func (s *student) Setname(name string) {
	s.Name = name
}
