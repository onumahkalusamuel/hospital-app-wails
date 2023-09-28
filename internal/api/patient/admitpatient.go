package patient

import (
	"net/http"
	"time"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

type AdmitPatientRequest struct {
	WardNumber string `json:"ward_number"`
	RoomNumber string `json:"room_number"`
	Note       string `json:"note"`
}

func AdmitPatient(c echo.Context) error {

	var req AdmitPatientRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request" + err.Error()})
	}

	if req.RoomNumber == "" {
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
		// Type:      "Admission",
		// Details: echo.Map{
		// 	"room_number": req.RoomNumber,
		// 	"ward_number": req.WardNumber,
		// 	"note":        req.Note,
		// },
		Date: &now,
	}

	// update patient record
	patient.CurrentStatus = "Admitted"
	patient.CurrentAppointment = "Emergency"
	config.DB.Updates(patient)

	if err := config.DB.Create(&patientHistory).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Patient admitted successfully."})
}
