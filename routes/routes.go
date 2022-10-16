package routes

import (
	"assigment2/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(ctl controller.Controller) error {
	r := gin.Default()

	r.GET("/orders", ctl.GetOrders)
	r.GET("/orders/:id")
	r.POST("/order", ctl.CreateOrder)
	r.PUT("/order/:id")
	r.DELETE("/order/:id")

	return r.Run("localhost:8080")
}
