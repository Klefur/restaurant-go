package routes

import (
	"github.com/gin-gonic/gin"
	controller "go-restaurant/controllers"
)

func InvoiceRoutes(router *gin.Engine) {
	
	router.GET("/invoices", controller.GetInvoices())
	router.GET("/invoices/:id", controller.GetInvoice())
	router.POST("/invoices", controller.CreateInvoice())
	router.PATCH("/invoices/:id", controller.UpdateInvoice())
}