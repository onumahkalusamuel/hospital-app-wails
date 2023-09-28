package payment

import (
	"fmt"
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {

	var payment models.Payment
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request: " + err.Error()})
	}

	if payment.AmountPaid == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "please enter all required fields."})
	}

	invoiceId := c.Param("invoice_id")

	// fetch the invoice
	var invoice models.Invoice
	invoice.ID = invoiceId
	if err := config.DB.Model(&models.Invoice{}).First(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invoice not found." + err.Error()})
	}

	// check if amount in balance is above or equal to paid amount
	if invoice.Balance < payment.AmountPaid {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "paid amount is above the remaining balance of " + fmt.Sprint(invoice.Balance)})
	}

	// create the payment
	payment.InvoiceID = invoiceId
	payment.StaffID = c.Get("ID").(string)
	payment.Balance = invoice.Balance - payment.AmountPaid

	if err := config.DB.Create(&payment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating record: " + err.Error()})
	}

	// save the new balance and completion status
	invoice.Balance = payment.Balance
	if invoice.Balance == 0 {
		invoice.Completed = 1
	}

	if err := config.DB.Save(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error updating balance: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"id": invoice.ID})
}
