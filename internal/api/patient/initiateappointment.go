package patient

import (
	"net/http"
	"time"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

type InitiateAppointmentRequest struct {
	AppointmentType string `json:"appointment_type"`
	Note            string `json:"note"`
}

func InitiateAppointment(c echo.Context) error {

	var req InitiateAppointmentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request" + err.Error()})
	}

	if req.AppointmentType == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "please enter all required fields."})
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
		// Type:      "Appointment",
		// Details: echo.Map{
		// 	"appointment_type": req.AppointmentType,
		// 	"note":             req.Note,
		// },
		Date: &now,
	}

	// update patient record
	patient.CurrentAppointment = req.AppointmentType
	config.DB.Updates(patient)

	if err := config.DB.Create(&patientHistory).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Appointment initated successfully."})
}
