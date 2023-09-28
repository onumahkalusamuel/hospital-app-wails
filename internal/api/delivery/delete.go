package delivery

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	delivery := &models.Delivery{}
	delivery.ID = c.Param("id")
	delivery.Read()

	if delivery.PatientID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	config.DB.Delete(delivery)

	return c.JSON(http.StatusOK, echo.Map{"message": "record deleted successfully"})
}
