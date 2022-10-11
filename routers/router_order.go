package routers

import (
	"assignment_2/controllers"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	router := gin.Default()

	router.GET("/orders", controllers.GetOrders)
	router.POST("/orders", controllers.CreateOrder)
	//MENYERAH UNTUK YG UPDATE DAN DELETE
	// router.PUT("/orders/:id_order", controllers.UpdateOrder)
	// router.DELETE("/orders/:id_order", controllers.UpdateOrder)

	return router
}
