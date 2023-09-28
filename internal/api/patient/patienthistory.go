package patient

import (
	"net/http"

	"hospital-app/internal/models"
	"hospital-app/pkg"

	"github.com/labstack/echo/v4"
)

func PatientHistory(c echo.Context) error {

	p := models.PatientHistory{PatientID: c.Param("patient_id")}
	pag, err := p.List(pkg.GrabFromContext(pkg.Pagination{}, c), &models.PatientHistory{PatientID: c.Param("patient_id")})
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, pag)
}
