package routers

import (
	"restapi/controllers"
	config "restapi/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func StartServer() *gin.Engine {

	// inisialisasi router
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()

	// Router API

	// post/create data
	router.POST("/orders", inDB.CreateOrders)

	// get all data
	router.GET("/orders", inDB.GetOrders)

	// update dataDeleteOrder
	router.PUT("/orders/:orderId", inDB.UpdateOrder)

	// DeleteOrder
	router.DELETE("/orders/:orderId", inDB.DeleteOrder)

	return router
}
