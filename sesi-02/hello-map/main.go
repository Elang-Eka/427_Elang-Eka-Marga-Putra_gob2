package main

import (
	"fmt"
)

func main() {

	//Map
	var person map[string]string

	person = map[string]string{}

	person["name"] = "Elang"
	person["age"] = "22"
	person["address"] = "Sri Minulyo"

	fmt.Println("name: ", person["name"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address: ", person["address"])

	fmt.Println("========================================================")
	//Map (Looping with map)
	var person2 = map[string]string{
		"name":    "Elang",
		"age":     "22",
		"address": "Sri Minulyo",
	}

	for key, value := range person2 {
		fmt.Println(key, " : ", value)
	}
	fmt.Println("========================================================")
	//Map (Deleting value)
	var person3 = map[string]string{
		"name":    "Elang",
		"age":     "22",
		"address": "Sri Minulyo",
	}

	fmt.Println("before deleting: ", person3)

	delete(person3, "age")

	fmt.Println("After Deleting: ", person3)
	fmt.Println("========================================================")
	//Map (Detecting a value)
	var person4 = map[string]string{
		"name":    "Elang",
		"age":     "22",
		"address": "Jalan Sri Minulyo",
	}

	value, exist := person4["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value does'nt exist")
	}

	delete(person4, "age")

	value, exist = person4["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}
	fmt.Println("========================================================")
	//Map (Combining slice with map)
	var people = []map[string]string{
		{"name": "Elang", "age": "22"},
		{"name": "Erie", "age": "21"},
		{"name": "Putra", "age": "23"},
	}
	for i, person5 := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person5["name"], person5["age"])
	}
	fmt.Println("========================================================")

}
