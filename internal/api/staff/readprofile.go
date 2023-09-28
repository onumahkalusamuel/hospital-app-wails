package staff

import (
	"net/http"

	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func ReadProfile(c echo.Context) error {
	staff := &models.Staff{}
	staff.ID = c.Get("ID").(string)
	staff.Read()
	if staff.Firstname == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	// staff.Password = ""
	return c.JSON(http.StatusOK, staff)
}
