package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetOrderItemById() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetOrderItemsByOrderId() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func itemsByOrderId(id string) (OrderItems []primitive.M, err error) {
	return nil, nil
}
