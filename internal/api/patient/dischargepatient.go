package patient

import (
	"net/http"
	"time"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

type DischargePatientRequest struct {
	Note string `json:"note"`
}

func DischargePatient(c echo.Context) error {

	var req DischargePatientRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request" + err.Error()})
	}

	patient := models.Patient{}
	patient.ID = c.Param("patient_id")
	patient.Read()
	if patient.Lastname == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "patient not found"})
	}

	now := time.Now()

	patientHistory := models.PatientHistory{
		StaffID:   c.Get("ID").(string),
		PatientID: patient.ID,
		// Type:      "Discharge",
		// Details: echo.Map{
		// 	"note": req.Note,
		// },
		Date: &now,
	}

	// update patient record
	patient.CurrentStatus = "Discharged"
	patient.CurrentAppointment = "Regular"
	config.DB.Updates(patient)

	if err := config.DB.Create(&patientHistory).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Patient discharged successfully."})
}
