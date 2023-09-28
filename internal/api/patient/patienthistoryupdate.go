package patient

import (
	"fmt"
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/helpers"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func PatientHistoryUpdate(c echo.Context) error {
	var err error
	var filename string
	var req PatientCreateParams
	var patientHistory models.PatientHistory

	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	// patientHistory.StaffID = c.Get("ID").(string)
	patientHistory.PatientID = c.Param("patient_id")
	patientHistory.ID = c.Param("history_id")
	patientHistory.Read()

	// subject
	patientHistory.Subject = req.Subject

	// admission
	patientHistory.Details.Admission.Note = req.AdmissionNote
	patientHistory.Details.Admission.RoomNumber = req.AdmissionRoomNumber
	patientHistory.Details.Admission.WardNumber = req.AdmissionWardNumber
	filename, err = helpers.SaveImage(c, "details[admission][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		if patientHistory.Details.Admission.Document != "" {
			helpers.DeleteImage(fmt.Sprintf("images_%v", patientHistory.PatientID), patientHistory.Details.Admission.Document)
		}
		patientHistory.Details.Admission.Document = filename
	}

	// appointment
	patientHistory.Details.Appointment.Note = req.AppointmentNote
	patientHistory.Details.Appointment.AppointmentType = req.AppointmentType
	filename, err = helpers.SaveImage(c, "details[appointment][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		if patientHistory.Details.Appointment.Document != "" {
			helpers.DeleteImage(fmt.Sprintf("images_%v", patientHistory.PatientID), patientHistory.Details.Appointment.Document)
		}
		patientHistory.Details.Appointment.Document = filename
	}

	// general
	patientHistory.Details.General.Note = req.GeneralNote
	filename, err = helpers.SaveImage(c, "details[general][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		if patientHistory.Details.General.Document != "" {
			helpers.DeleteImage(fmt.Sprintf("images_%v", patientHistory.PatientID), patientHistory.Details.General.Document)
		}
		patientHistory.Details.General.Document = filename
	}

	// discharge
	patientHistory.Details.Discharge.Note = req.DischargeNote
	filename, err = helpers.SaveImage(c, "details[discharge][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		if patientHistory.Details.Discharge.Document != "" {
			helpers.DeleteImage(fmt.Sprintf("images_%v", patientHistory.PatientID), patientHistory.Details.Discharge.Document)
		}
		patientHistory.Details.Discharge.Document = filename
	}

	// diagnosis
	patientHistory.Details.Diagnosis.Note = req.DiagnosisNote
	filename, err = helpers.SaveImage(c, "details[diagnosis][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		if patientHistory.Details.Diagnosis.Document != "" {
			helpers.DeleteImage(fmt.Sprintf("images_%v", patientHistory.PatientID), patientHistory.Details.Diagnosis.Document)
		}
		patientHistory.Details.Diagnosis.Document = filename
	}

	// examination
	patientHistory.Details.Examination.Note = req.ExaminationNote
	filename, err = helpers.SaveImage(c, "details[examination][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		if patientHistory.Details.Examination.Document != "" {
			helpers.DeleteImage(fmt.Sprintf("images_%v", patientHistory.PatientID), patientHistory.Details.Examination.Document)
		}
		patientHistory.Details.Examination.Document = filename
	}

	// testresult
	patientHistory.Details.TestResult.Note = req.TestResultNote
	filename, err = helpers.SaveImage(c, "details[testresult][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		if patientHistory.Details.TestResult.Document != "" {
			helpers.DeleteImage(fmt.Sprintf("images_%v", patientHistory.PatientID), patientHistory.Details.TestResult.Document)
		}
		patientHistory.Details.TestResult.Document = filename
	}

	// treatment
	patientHistory.Details.Treatment.Note = req.TreatmentNote
	filename, err = helpers.SaveImage(c, "details[treatment][document]", fmt.Sprintf("images_%v", patientHistory.PatientID))
	if err == nil {
		if patientHistory.Details.Treatment.Document != "" {
			helpers.DeleteImage(fmt.Sprintf("images_%v", patientHistory.PatientID), patientHistory.Details.Treatment.Document)
		}
		patientHistory.Details.Treatment.Document = filename
	}

	// update
	config.DB.Updates(patientHistory)

	return c.JSON(http.StatusOK, echo.Map{"message": "record updated successfully."})
}
