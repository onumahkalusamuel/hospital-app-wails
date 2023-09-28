package staff

import (
	"net/http"

	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func ReadOne(c echo.Context) error {
	staff := &models.Staff{}
	staff.ID = c.Param("id")
	staff.Read()
	if staff.Username == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Staff record not found"})
	}

	staff.Password = ""
	return c.JSON(http.StatusOK, staff)
}
