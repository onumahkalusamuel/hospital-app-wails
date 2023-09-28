package patient

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func ReadOne(c echo.Context) error {
	var patient models.Patient
	patient.ID = c.Param("id")
	config.DB.Model(&models.Patient{}).First(&patient)
	if patient.Firstname == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "account not found"})
	}

	return c.JSON(http.StatusOK, patient)
}
