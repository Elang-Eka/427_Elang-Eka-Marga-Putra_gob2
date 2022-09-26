package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Goroutines
	go goroutine()
	fmt.Println("============================================")

	// Goroutines (Asynchronous process)
	fmt.Println("main execution started")

	go firstProcess(8)

	secondProcess(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	fmt.Println("main execution ended")
	fmt.Println("============================================")
	// Goroutines (Asynchronous process) time second
	fmt.Println("main execution started")

	go firstProcess(8)

	secondProcess(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	time.Sleep(time.Second * 2)
	fmt.Println("main execution ended")
	fmt.Println("============================================")

}

// function goroutines
func goroutine() {
	fmt.Println("Hello")
}

// function firstProcess
func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

// function secondProcess
func secondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}
