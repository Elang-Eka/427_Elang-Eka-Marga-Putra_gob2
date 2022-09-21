package main

import "fmt"

func main() {
	//aliase 1
	var a uint8 = 10
	var b byte

	b = a
	_ = b

	//aliase 2
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type : %T\n", hour)
}
