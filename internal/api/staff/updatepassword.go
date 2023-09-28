package staff

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"
	"hospital-app/pkg"

	"github.com/labstack/echo/v4"
)

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func UpdatePassword(c echo.Context) error {

	var req UpdatePasswordRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request: " + err.Error()})
	}

	staff := &models.Staff{}
	staff.ID = c.Get("id").(string)

	if err := staff.Read(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request: " + err.Error()})
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "enter old and new password."})
	}

	if pkg.CheckPassword(staff.Password, req.OldPassword) == false {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "wrong old password provided."})
	}

	staff.Password = pkg.HashPassword(req.NewPassword)
	config.DB.Model(&models.Staff{}).Updates(staff)

	return c.JSON(http.StatusOK, echo.Map{"message": "password changed successfully."})
}
