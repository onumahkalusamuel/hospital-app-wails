package delivery

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	var delivery models.Delivery

	if err := c.Bind(&delivery); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	delivery.ID = c.Param("id")

	config.DB.Updates(delivery)

	return c.JSON(http.StatusOK, echo.Map{"message": "Record updated successfully."})
}
