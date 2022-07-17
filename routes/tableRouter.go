package routes

import (
	"github.com/OmkarMahajan11/RestaurantManagement-backend/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetAllTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTableById())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())
}
