package patient

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

	patient := &models.Patient{}
	patient.ID = c.Param("id")
	patient.Read()

	if patient.Firstname == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "account not found"})
	}

	config.DB.Delete(patient)
	// patient.Delete()

	return c.JSON(http.StatusOK, echo.Map{"message": "account deleted successfully"})
}
