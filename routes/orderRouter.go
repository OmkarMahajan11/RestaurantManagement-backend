package routes

import (
	"github.com/OmkarMahajan11/RestaurantManagement-backend/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controllers.GetAllOrders())
	incomingRoutes.GET("/orders/:order_id", controllers.GetOrderById())
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controllers.UpdateOrder())
}
