package routes

import (
	"github.com/OmkarMahajan11/RestaurantManagement-backend/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controllers.GetAllInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoiceById())
	incomingRoutes.POST("/invoices", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}
