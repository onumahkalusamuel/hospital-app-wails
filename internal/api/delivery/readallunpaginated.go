package delivery

import (
	"net/http"
	"time"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

type ReadAllRequest struct {
	Timeline string    `json:"timeline"`
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}

func ReadAllUnpaginated(c echo.Context) error {
	var req ReadAllRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	db := config.DB.Model(&models.Delivery{})

	// calculate what this year value is
	if req.Timeline == "thisyear" {
		req.DateFrom = time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Now().Location())
		req.DateTo = time.Date(time.Now().Year(), 12, 31, 23, 59, 59, 0, time.Now().Location())
	}

	// add the between clause for != alltime
	if req.Timeline != "alltime" {
		db.Where("delivery_date_time between ? and ?", req.DateFrom, req.DateTo)
	}

	var deliveries []*models.Delivery

	db.Preload("Patient").Find(&deliveries)

	return c.JSON(http.StatusOK, deliveries)
}
