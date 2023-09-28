package staff

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"
	"hospital-app/pkg"

	"github.com/labstack/echo/v4"
)

type UpdateStaffRequest struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Role       uint   `json:"role"`
}

func Update(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	var req UpdateStaffRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	staff := &models.Staff{}
	staff.ID = c.Param("id")
	staff.Read()

	if err := staff.Read(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request: " + err.Error()})
	}

	if req.Firstname != staff.Firstname {
		staff.Firstname = req.Firstname
	}

	if req.Lastname != staff.Lastname {
		staff.Lastname = req.Lastname
	}

	if req.Middlename != staff.Middlename {
		staff.Middlename = req.Middlename
	}

	if req.Role != staff.Role {
		staff.Role = req.Role
	}

	if req.Password != "" {
		staff.UpdateSingle("password", pkg.HashPassword(req.Password))
	}

	config.DB.Updates(&staff)

	return c.JSON(http.StatusOK, echo.Map{"message": "record updated successfully."})
}
