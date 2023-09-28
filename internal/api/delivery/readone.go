package delivery

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func ReadOne(c echo.Context) error {
	var delivery models.Delivery
	delivery.ID = c.Param("id")
	config.DB.Model(&models.Delivery{}).Preload("Patient").First(&delivery)
	if delivery.PatientID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	return c.JSON(http.StatusOK, delivery)
}
