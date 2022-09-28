package main

import (
	"fmt"
	"os"
)

func main() {
	// Exit
	defer fmt.Println("Invoke with defer")

	fmt.Println("Before exiting")
	fmt.Println("======================================")
	os.Exit(1)
	// ketika exit maka bagian ini tidak akan ditampilkan
	fmt.Println("Tidak tampil")
}
