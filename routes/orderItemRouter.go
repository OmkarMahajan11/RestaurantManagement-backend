package routes

import (
	"github.com/OmkarMahajan11/RestaurantManagement-backend/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetAllOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItemById())
	incomingRoutes.GET("/orderItems/:order_id", controllers.GetOrderItemsByOrderId())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
