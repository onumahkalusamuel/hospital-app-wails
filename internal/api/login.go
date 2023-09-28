package api

import (
	"net/http"
	"time"

	"hospital-app/config"
	"hospital-app/internal/models"
	"hospital-app/pkg"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	ID       string `json:"id"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Role     uint   `json:"role"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var loginRequest LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	staff := models.Staff{
		Username: loginRequest.Username,
	}

	staff.Read()

	if staff.ID == "" || !pkg.CheckPassword(staff.Password, loginRequest.Password) {
		return c.JSON(403, echo.Map{"message": "Invalid username or password"})
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		staff.ID, staff.Lastname, staff.Username, staff.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return err
	}
	return c.JSON(200, echo.Map{
		"message":        "Logged in successfully",
		"jwt":            t,
		"asset_base_url": "//" + c.Request().Host + "/",
	})
}
