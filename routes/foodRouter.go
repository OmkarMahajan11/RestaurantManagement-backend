package routes

import (
	"github.com/OmkarMahajan11/RestaurantManagement-backend/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetAllFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFoodById())
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controllers.UpdateFood())
}
