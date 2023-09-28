package api

import (
	"net/http"

	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func HospitalDetails(c echo.Context) error {

	hospital := echo.Map{
		"hospital_name":    "",
		"hospital_address": "",
		"hospital_email":   "",
		"hospital_phone":   "",
		"hospital_logo":    "",
		"asset_base_url":   "",
	}
	var setting *models.Settings

	// name
	setting = &models.Settings{Setting: "hospital_name"}
	setting.Read()
	hospital["hospital_name"] = setting.Value

	// address
	setting = &models.Settings{Setting: "hospital_address"}
	setting.Read()
	hospital["hospital_address"] = setting.Value

	// email
	setting = &models.Settings{Setting: "hospital_email"}
	setting.Read()
	hospital["hospital_email"] = setting.Value

	// phone
	setting = &models.Settings{Setting: "hospital_phone"}
	setting.Read()
	hospital["hospital_phone"] = setting.Value

	// logo
	setting = &models.Settings{Setting: "hospital_logo"}
	setting.Read()
	hospital["hospital_logo"] = setting.Value

	hospital["asset_base_url"] = "//" + c.Request().Host

	// return
	return c.JSON(http.StatusOK, hospital)
}
