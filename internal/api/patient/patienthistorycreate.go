package patient

import (
	"fmt"
	"net/http"
	"time"

	"hospital-app/config"
	"hospital-app/internal/helpers"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

type PatientCreateParams struct {
	PatientID           string `form:"patient_id"`
	Subject             string `form:"subject"`
	AdmissionNote       string `form:"details[admission][note]"`
	AdmissionDocument   string `form:"details[admission][document]"`
	AdmissionWardNumber string `form:"details[admission][ward_number]"`
	AdmissionRoomNumber string `form:"details[admission][room_number]"`
	AppointmentNote     string `form:"details[appointment][note]"`
	AppointmentDocument string `form:"details[appointment][document]"`
	AppointmentType     string `form:"details[appointment][appointment_type]"`
	GeneralNote         string `form:"details[general][note]"`
	GeneralDocument     string `form:"details[general][document]"`
	DischargeNote       string `form:"details[discharge][note]"`
	DischargeDocument   string `form:"details[discharge][document]"`
	DiagnosisNote       string `form:"details[diagnosis][note]"`
	DiagnosisDocument   string `form:"details[diagnosis][document]"`
	ExaminationNote     string `form:"details[examination][note]"`
	ExaminationDocument string `form:"details[examination][document]"`
	TestResultNote      string `form:"details[testresult][note]"`
	TestResultDocument  string `form:"details[testresult][document]"`
	TreatmentNote       string `form:"details[treatment][note]"`
	TreatmentDocument   string `form:"details[treatment][document]"`
}

func PatientHistoryCreate(c echo.Context) error {

	var err error
	var filename string
	var req PatientCreateParams
	var patientHistory models.PatientHistory

	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	now := time.Now()
	patientHistory.StaffID = c.Get("ID").(string)
	patientHistory.PatientID = c.Param("patient_id")
	patientHistory.Date = &now

	// subject
	patientHistory.Subject = req.Subject

	// admission
	patientHistory.Details.Admission.Note = req.AdmissionNote
	patientHistory.Details.Admission.RoomNumber = req.AdmissionRoomNumber
	patientHistory.Details.Admission.WardNumber = req.AdmissionWardNumber
	filename, err = helpers.SaveImage(c, "details[admission][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		patientHistory.Details.Admission.Document = filename
	}

	// appointment
	patientHistory.Details.Appointment.Note = req.AppointmentNote
	patientHistory.Details.Appointment.AppointmentType = req.AppointmentType
	filename, err = helpers.SaveImage(c, "details[appointment][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		patientHistory.Details.Appointment.Document = filename
	}

	// general
	patientHistory.Details.General.Note = req.GeneralNote
	filename, err = helpers.SaveImage(c, "details[general][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		patientHistory.Details.General.Document = filename
	}

	// discharge
	patientHistory.Details.Discharge.Note = req.DischargeNote
	filename, err = helpers.SaveImage(c, "details[discharge][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		patientHistory.Details.Discharge.Document = filename
	}

	// diagnosis
	patientHistory.Details.Diagnosis.Note = req.DiagnosisNote
	filename, err = helpers.SaveImage(c, "details[diagnosis][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		patientHistory.Details.Diagnosis.Document = filename
	}

	// examination
	patientHistory.Details.Examination.Note = req.ExaminationNote
	filename, err = helpers.SaveImage(c, "details[examination][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		patientHistory.Details.Examination.Document = filename
	}

	// testresult
	patientHistory.Details.TestResult.Note = req.TestResultNote
	filename, err = helpers.SaveImage(c, "details[testresult][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		patientHistory.Details.TestResult.Document = filename
	}

	// treatment
	patientHistory.Details.Treatment.Note = req.TreatmentNote
	filename, err = helpers.SaveImage(c, "details[treatment][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		patientHistory.Details.Treatment.Document = filename
	}

	if err := config.DB.Create(&patientHistory).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"id": patientHistory.ID})
}
