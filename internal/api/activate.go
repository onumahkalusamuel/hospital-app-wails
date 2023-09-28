package api

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"
	"hospital-app/pkg"

	"github.com/labstack/echo/v4"
)

type ActivateRequest struct {
	ActivationCode string `json:"activation_code"`
}

func Activate(c echo.Context) error {
	// edit later to tie hosiptal name to activation code
	var activateRequest ActivateRequest

	err := c.Bind(&activateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	installationCode := config.HashedID
	activationCode := activateRequest.ActivationCode
	checkActivationCode := pkg.GenActivationCode(installationCode)

	if len(installationCode) < 10 || len(activationCode) != 14 || activationCode != checkActivationCode {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid activation code",
		})
	}

	//update settings
	settings := &models.Settings{Setting: "activation_code"}
	settings.Read()
	settings.UpdateSingle("value", activationCode)

	// set activation in config
	config.ActivationCode = activationCode

	return c.JSON(http.StatusOK, echo.Map{
		"message": "System activated successfully.",
	})
}
