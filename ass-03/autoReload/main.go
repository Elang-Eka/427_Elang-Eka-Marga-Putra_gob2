package main

import (
	"autoReload/controllers"
	config "autoReload/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var PORT = ":8080"
	// inisialisasi router
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()
	// router API
	router.POST("/autoReload/", inDB.PostData)
	router.PUT("/autoReload/:id", inDB.UpdateData)
	router.GET("/autoReload/:id", inDB.Getdata)
	router.GET("/autoReload", inDB.Getdatas)
	router.Run(PORT)
}
