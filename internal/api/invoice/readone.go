package invoice

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func ReadOne(c echo.Context) error {
	invoice := &models.Invoice{}
	invoice.ID = c.Param("id")
	config.DB.Model(&models.Invoice{}).Preload("Payments").Preload("Patient").First(&invoice)

	if invoice.PatientID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	return c.JSON(http.StatusOK, invoice)
}
