package main

import (
	"hello/helpers"
)

func main() {
	// Exported & Unexported

	helpers.Greet()

	var person = helpers.Person{}

	person.Invokegreet()

	// Init function
}
