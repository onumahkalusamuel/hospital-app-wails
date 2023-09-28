package patient

import (
	"net/http"

	"hospital-app/internal/models"
	"hospital-app/pkg"

	"github.com/labstack/echo/v4"
)

func ReadAll(c echo.Context) error {
	p := models.Patient{}
	pag, err := p.List(pkg.GrabFromContext(pkg.Pagination{}, c), []string{})
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, pag)
}
