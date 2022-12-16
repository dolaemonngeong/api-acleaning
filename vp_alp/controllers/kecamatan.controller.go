package controllers

import (
	"net/http"
	"vp_alp/models"

	"github.com/labstack/echo/v4"
)

func FetchAllKecamatan(c echo.Context) error {

	result, err := models.FetchAllKecamatan()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}