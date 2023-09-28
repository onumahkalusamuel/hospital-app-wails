package delivery

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {

	var delivery models.Delivery
	if err := c.Bind(&delivery); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request " + err.Error()})
	}

	if delivery.BabySex == "" || delivery.PatientID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "please enter all required fields."})
	}

	delivery.StaffID = c.Get("ID").(string)

	if err := config.DB.Create(&delivery).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"id": delivery.ID})
}
