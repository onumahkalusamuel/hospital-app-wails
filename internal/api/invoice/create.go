package invoice

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/helpers"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {

	var invoice models.Invoice
	if err := c.Bind(&invoice); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request" + err.Error()})
	}

	if invoice.PatientID == "" || len(invoice.Details) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "please enter all required fields."})
	}

	invoice.Subtotal = 0
	for x, details := range invoice.Details {
		invoice.Details[x].Amount = details.Qty * details.Price
		invoice.Subtotal += invoice.Details[x].Amount
	}

	invoice.StaffID = c.Get("ID").(string)
	invoice.Balance = invoice.Subtotal - invoice.Discount
	invoice.Amount = invoice.Balance

	if invoice.InvoiceNumber == "" {
		invoice.InvoiceNumber = helpers.GetRandKey(7)
	}

	if err := config.DB.Create(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"id": invoice.ID})
}
