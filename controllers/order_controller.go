package controllers

import (
	"assignment_2/database"
	"assignment_2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var listOrder = []models.Orders{}

func GetOrders(ctx *gin.Context) {
	db := database.GetDB()
	db.Preload("Items").Find(&listOrder)
	ctx.JSON(http.StatusOK, listOrder)
}

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Orders
	db := database.GetDB()

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Wrong Params")
		return
	}
	db.Create(&newOrder)
	ctx.JSON(http.StatusOK, newOrder)
}

// func UpdateOrder(ctx *gin.Context) {
// 	var updateOrder models.Orders

// 	orderId := ctx.Params("")

// }

//Menyerah untuk API delete dan update
