package controller

import (
	"assigment2/database"
	"assigment2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database.Database
}

func New(db database.Database) Controller {
	return Controller{
		db: db,
	}
}

func (c Controller) GetOrders(ctx *gin.Context) {
	orders, err := c.db.GetOrders()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error get data",
		})
	}
	ctx.JSON(http.StatusOK, orders)
}

func (c Controller) CreateOrder(ctx *gin.Context) {
	var newOrder model.Order
	err := ctx.BindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error bind json request",
		})
	}

	orderResult, err := c.db.CreateOrder(newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error create order",
		})
	}

	ctx.JSON(http.StatusOK, orderResult)
}
