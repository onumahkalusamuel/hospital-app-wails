package api

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"
	"hospital-app/pkg"

	"github.com/labstack/echo/v4"
)

type CreateAdminRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func CreateAdmin(c echo.Context) error {
	var createAdminRequest CreateAdminRequest

	err := c.Bind(&createAdminRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	checkadmin := &models.Staff{Role: config.ADMIN_ROLE}
	checkadmin.Read()

	if checkadmin.ID != "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Admin user already exist.",
		})
	}

	if createAdminRequest.Firstname == "" || createAdminRequest.Lastname == "" || createAdminRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    5,
			"message": "Please provide valid details",
		})
	}

	//continue
	admin := &models.Staff{
		Firstname: createAdminRequest.Firstname,
		Lastname:  createAdminRequest.Lastname,
		Username:  createAdminRequest.Username,
		Password:  pkg.HashPassword(createAdminRequest.Password),
		Role:      config.ADMIN_ROLE,
	}

	if err := admin.Create(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    5,
			"message": "Error occured: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Admin created successfully."})
}
