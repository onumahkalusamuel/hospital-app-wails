package patient

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func PatientHistoryDelete(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	patientHistory := &models.PatientHistory{}
	patientHistory.ID = c.Param("history_id")
	patientHistory.PatientID = c.Param("patient_id")
	patientHistory.Read()

	if patientHistory.StaffID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	config.DB.Delete(patientHistory)

	return c.JSON(http.StatusOK, echo.Map{"message": "record deleted successfully"})
}
