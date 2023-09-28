package delivery

import (
	"net/http"

	"hospital-app/internal/models"
	"hospital-app/pkg"

	"github.com/labstack/echo/v4"
)

func ReadAll(c echo.Context) error {
	// var deliveries []models.Delivery
	// config.DB.Model(&models.Delivery{}).Preload("Patient").Find(&deliveries)
	// return c.JSON(http.StatusOK, deliveries)

	handle := models.Delivery{}
	if patient_id := c.QueryParam("patient_id"); patient_id != "" {
		handle.PatientID = patient_id
	}

	if baby_sex := c.QueryParam("baby_sex"); baby_sex != "" {
		handle.BabySex = baby_sex
	}

	if condition := c.QueryParam("condition"); condition != "" {
		handle.Condition = condition
	}

	if delivery_mode := c.QueryParam("delivery_mode"); delivery_mode != "" {
		handle.DeliveryMode = delivery_mode
	}

	list, err := handle.List(pkg.GrabFromContext(pkg.Pagination{}, c), []string{"Patient"}, &handle)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, list)
}
