package controllers

import (
	"assignment-2/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrders(c *gin.Context) {
	var newOrder models.Order

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newOrder = QueryCreate(newOrder)

	c.JSON(http.StatusCreated, gin.H{
		"data":    newOrder,
		"message": "Data sucessfully created",
		"status":  http.StatusCreated,
	})
}

func GetAllOrders(c *gin.Context) {
	orders := QueryGetAll()

	c.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"message": "Orders fetched sucessfully",
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})

}

func DeleteOrder(c *gin.Context) {
	orderID := c.Param("orderID")

	convertedOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Invalid Params"),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	QueryDeleteByID(uint(convertedOrderID))

	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": fmt.Sprintf("Order with ID %v Has been sucessfully deleted", orderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}

func UpdateOrderByID(c *gin.Context) {
	var updatedOrder models.Order
	orderID := c.Param("orderID")

	convertedOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Invalid Params"),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedOrder = QueryUpdateByID(updatedOrder, uint(convertedOrderID))

	c.JSON(http.StatusOK, gin.H{
		"data":    updatedOrder,
		"message": fmt.Sprintf("Order with ID %v Has been sucessfully updated", orderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})

}
