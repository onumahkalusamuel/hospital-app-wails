package api

import (
	"net/http"

	"hospital-app/internal/helpers"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

type HospitalSetupRequest struct {
	HospitalName    string `json:"hospital_name"`
	HospitalAddress string `json:"hospital_address"`
	HospitalEmail   string `json:"hospital_email"`
	HospitalPhone   string `json:"hospital_phone"`
	HospitalLogo    string `json:"hospital_logo"`
}

func HospitalSetup(c echo.Context) error {
	var hospitalSetupRequest HospitalSetupRequest

	hospitalSetupRequest.HospitalName = c.FormValue("hospital_name")
	hospitalSetupRequest.HospitalAddress = c.FormValue("hospital_address")
	hospitalSetupRequest.HospitalEmail = c.FormValue("hospital_email")
	hospitalSetupRequest.HospitalPhone = c.FormValue("hospital_phone")

	checkHospital := &models.Settings{Setting: "hospital_name"}
	checkHospital.Read()

	if checkHospital.Value != "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Hospital record already setup.",
		})
	}

	if hospitalSetupRequest.HospitalName == "" || hospitalSetupRequest.HospitalAddress == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Please provide valid details - 2",
		})
	}

	var setting *models.Settings
	// name
	setting = &models.Settings{Setting: "hospital_name"}
	setting.Read()
	setting.UpdateSingle("Value", hospitalSetupRequest.HospitalName)

	// address
	setting = &models.Settings{Setting: "hospital_address"}
	setting.Read()
	setting.UpdateSingle("Value", hospitalSetupRequest.HospitalAddress)

	// email
	setting = &models.Settings{Setting: "hospital_email"}
	setting.Read()
	setting.UpdateSingle("Value", hospitalSetupRequest.HospitalEmail)

	// phone
	setting = &models.Settings{Setting: "hospital_phone"}
	setting.Read()
	setting.UpdateSingle("Value", hospitalSetupRequest.HospitalPhone)

	filename, err := helpers.SaveImage(c, "hospital_logo", "logo")

	if err == nil {
		setting = &models.Settings{Setting: "hospital_logo"}
		setting.Read()
		setting.UpdateSingle("Value", filename)
	}

	// return
	return c.JSON(http.StatusOK, echo.Map{"message": "Hospital record updated successfully."})
}
