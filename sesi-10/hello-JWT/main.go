package main

import (
	"hello/database"
	"hello/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
