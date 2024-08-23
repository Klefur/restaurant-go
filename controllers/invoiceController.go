package controllers

import (
	"go-restaurant/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invoices []models.Invoice

		err := db.Find(&invoices).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occurred while fetching invoices"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": invoices})
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

		invoiceId := c.Param("id")
		var invoice models.Invoice

		err := db.Find(&invoice, invoiceId).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occurred while fetching invoice"})
			return
		}

		if invoice.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "invoice was not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": invoice})
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

		var invoice models.Invoice

		err := c.BindJSON(&invoice)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		var order models.Order

		err = db.First(&order, invoice.Order_id).Error
		if err != nil || order.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "menu was not found"})
			return
		}

		if invoice.Payment_status == "" {
			status := "PENDING"
			invoice.Payment_status = status
		}

		err = db.Create(&invoice).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occurred while creating invoice"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": invoice})
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invoice models.Invoice

		err := c.BindJSON(&invoice)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		invoiceId := c.Param("id")
		err = db.Model(&invoice).Where("id = ?", invoiceId).Updates(&invoice).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occurred while updating invoice"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": invoice})
	}
}
