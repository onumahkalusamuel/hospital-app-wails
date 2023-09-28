package staff

import (
	"fmt"
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"
	"hospital-app/pkg"

	"github.com/labstack/echo/v4"
)

type UpdateProfileRequest struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

func UpdateProfile(c echo.Context) error {
	var req UpdateProfileRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	fmt.Println(req)

	staff := &models.Staff{}
	staff.ID = c.Get("ID").(string)

	if err := staff.Read(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request: " + err.Error()})
	}

	staff.Firstname = req.Firstname
	staff.Lastname = req.Lastname
	staff.Middlename = req.Middlename
	staff.Phone = req.Phone
	staff.Email = req.Email

	if req.Password != "" {
		staff.UpdateSingle("password", pkg.HashPassword(req.Password))
	}

	config.DB.Updates(&staff)

	return c.JSON(http.StatusOK, echo.Map{"message": "profile updated successfully."})
}
