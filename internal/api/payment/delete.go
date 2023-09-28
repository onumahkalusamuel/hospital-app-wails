package payment

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	invoiceId := c.Param("invoice_id")
	// fetch the invoice
	var invoice models.Invoice
	invoice.ID = invoiceId
	if err := config.DB.Model(&models.Invoice{}).First(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invoice not found." + err.Error()})
	}

	payment := &models.Payment{}
	payment.ID = c.Param("id")
	payment.InvoiceID = invoiceId
	payment.Read()

	if payment.InvoiceID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	// save the new balance and completion status
	invoice.Balance = invoice.Balance + payment.AmountPaid
	invoice.Completed = 0

	if err := config.DB.Save(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error updating balance: " + err.Error()})
	}

	// then delete the payment
	config.DB.Delete(payment)

	return c.JSON(http.StatusOK, echo.Map{"message": "record deleted successfully"})
}
