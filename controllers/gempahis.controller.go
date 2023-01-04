package controllers

import (
	"net/http"
	"api_gempa/models"
	"github.com/labstack/echo/v4"
)

//get all gempa
func FetchAllGempaHis(c echo.Context) error{
	res, err := models.GetGempaHis()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

//store all gempa
func StoreGempaHis(c echo.Context) error{
	wilayah := c.FormValue("wilayah")
	tanggal := c.FormValue("tanggal")
	magnitudo := c.FormValue("magnitudo")

	res, err := models.StoreGempaHis(wilayah,tanggal,magnitudo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}