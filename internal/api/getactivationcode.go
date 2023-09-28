package api

import (
	"net/http"

	"hospital-app/pkg"

	"github.com/labstack/echo/v4"
)

type ActivationRequest struct {
	InstallationCode string `json:"installation_code"`
}

func GetActivationCode(c echo.Context) error {
	// edit later to tie hosiptal name to activation code

	var activationRequest ActivationRequest
	err := c.Bind(&activationRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	if len(activationRequest.InstallationCode) < 10 {
		return c.JSON(http.StatusOK, echo.Map{
			"code":    3,
			"message": "Invalid installation code",
		})
	}

	activationCode := pkg.GenActivationCode(activationRequest.InstallationCode)

	return c.JSON(http.StatusOK, echo.Map{
		"activation_code": activationCode,
	})
}
