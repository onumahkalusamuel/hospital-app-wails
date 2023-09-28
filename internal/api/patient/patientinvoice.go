package patient

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func PatientInvoice(c echo.Context) error {
	patient := &models.Patient{}
	patient.ID = c.Param("patient_id")
	config.DB.Model(&models.Patient{}).Preload("Invoices").First(&patient)

	if patient.CardNo == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	return c.JSON(http.StatusOK, patient)
}
