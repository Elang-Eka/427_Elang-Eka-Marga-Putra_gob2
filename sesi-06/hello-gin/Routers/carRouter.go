package Routers

import (
	controllers "hello/Controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	// inisialisasi router
	router := gin.Default()

	// router API
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)
	router.GET("/getAllCars", controllers.GetCars)

	return router
}
