package patient

import (
	"fmt"
	"net/http"

	"hospital-app/config"
	"hospital-app/internal/models"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {

	var patient models.Patient
	if err := c.Bind(&patient); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request" + err.Error()})
	}

	if patient.Firstname == "" || patient.Lastname == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "please enter all required fields."})
	}

	// check cardno
	if patient.CardNo != "" {
		var inuse = []models.Patient{}
		config.DB.Unscoped().Where("card_no='" + patient.CardNo + "'").Find(&inuse)

		if len(inuse) > 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": fmt.Sprintf("Card No already in use by %s %s", inuse[0].Firstname, inuse[0].Lastname)})
		}
	}

	patient.StaffID = c.Get("ID").(string)

	if err := config.DB.Create(&patient).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating account: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"id": patient.ID})
}
