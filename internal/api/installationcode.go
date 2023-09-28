package api

import (
	"net/http"

	"hospital-app/config"

	"github.com/labstack/echo/v4"
)

func InstallationCode(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"installation_code": config.HashedID})
}
