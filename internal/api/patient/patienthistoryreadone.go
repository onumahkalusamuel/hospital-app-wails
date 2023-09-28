package patient

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func PatientHistoryReadOne(c echo.Context) error {
	var patientHistory models.PatientHistory
	patientHistory.ID = c.Param("history_id")
	patientHistory.PatientID = c.Param("patient_id")
	config.DB.Model(&models.Patient{}).Preload("Patient").First(&patientHistory)
	if patientHistory.StaffID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	return c.JSON(http.StatusOK, patientHistory)
}
