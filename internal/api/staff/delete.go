package staff

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

	staff := &models.Staff{}
	staff.ID = c.Param("id")
	staff.Read()

	if staff.Role == config.ADMIN_ROLE {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Cannot delete the Super Admin"})
	}

	if staff.Username == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "account not found"})
	}

	staff.Delete()

	return c.JSON(http.StatusOK, echo.Map{"message": "account deleted successfully"})
}
