package staff

import (
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"
	"hospital-app/pkg"

	"github.com/labstack/echo/v4"
)

type CreateStaffRequest struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Role       uint   `json:"role"`
	Sex        string `json:"sex"`
}

func Create(c echo.Context) error {

	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	var staff CreateStaffRequest
	if err := c.Bind(&staff); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request" + err.Error()})
	}

	if staff.Firstname == "" || staff.Lastname == "" || staff.Username == "" || staff.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "please enter all required fields."})
	}

	// check username
	var inuse = []models.Staff{}
	config.DB.Unscoped().Where("username LIKE ?", staff.Username).Find(&inuse)

	if len(inuse) > 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "username already in use"})
	}

	// hash password
	staff.Password = pkg.HashPassword(staff.Password)

	newstaff := models.Staff{
		Firstname:  staff.Firstname,
		Middlename: staff.Middlename,
		Lastname:   staff.Lastname,
		Username:   staff.Username,
		Email:      staff.Email,
		Phone:      staff.Phone,
		Password:   staff.Password,
		Role:       staff.Role,
		Sex:        staff.Sex,
	}

	if err := config.DB.Create(&newstaff).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating account: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"id": newstaff.ID})
}
