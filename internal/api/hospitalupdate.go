package api

import (
	"net/http"

	"hospital-app/internal/helpers"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

type HospitalUpdateRequest struct {
	HospitalName    string `json:"hospital_name"`
	HospitalAddress string `json:"hospital_address"`
	HospitalEmail   string `json:"hospital_email"`
	HospitalPhone   string `json:"hospital_phone"`
	HospitalLogo    string `json:"hospital_logo"`
}

func HospitalUpdate(c echo.Context) error {
	var updateRequest HospitalUpdateRequest

	// if err := c.Bind(&updateRequest); err != nil {
	// 	return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	// }

	updateRequest.HospitalName = c.FormValue("hospital_name")
	updateRequest.HospitalAddress = c.FormValue("hospital_address")
	updateRequest.HospitalEmail = c.FormValue("hospital_email")
	updateRequest.HospitalPhone = c.FormValue("hospital_phone")

	if updateRequest.HospitalName == "" || updateRequest.HospitalAddress == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Please provide valid details",
		})
	}

	var setting *models.Settings
	// name
	setting = &models.Settings{Setting: "hospital_name"}
	setting.Read()
	if setting.Value != updateRequest.HospitalName {
		setting.UpdateSingle("Value", updateRequest.HospitalName)
	}

	// address
	setting = &models.Settings{Setting: "hospital_address"}
	setting.Read()
	if setting.Value != updateRequest.HospitalAddress {
		setting.UpdateSingle("Value", updateRequest.HospitalAddress)
	}

	// email
	setting = &models.Settings{Setting: "hospital_email"}
	setting.Read()
	if setting.Value != updateRequest.HospitalEmail {
		setting.UpdateSingle("Value", updateRequest.HospitalEmail)
	}

	// phone
	setting = &models.Settings{Setting: "hospital_phone"}
	setting.Read()
	if setting.Value != updateRequest.HospitalPhone {
		setting.UpdateSingle("Value", updateRequest.HospitalPhone)
	}

	// image
	oldLogo := &models.Settings{Setting: "hospital_logo"}
	oldLogo.Read()

	filename, err := helpers.SaveImage(c, "hospital_logo", "logo")

	if err == nil {
		setting = &models.Settings{Setting: "hospital_logo"}
		setting.Read()
		setting.UpdateSingle("Value", filename)
		if oldLogo.Value != "" {
			helpers.DeleteImage("logo", oldLogo.Value)
		}
	}

	// return
	return c.JSON(http.StatusOK, echo.Map{"message": "Hospital record updated successfully."})
}
