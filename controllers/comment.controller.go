package controllers

import(
	"net/http"
	"api_gempa/models"
	"github.com/labstack/echo/v4"
)

func FetchAllComment(c echo.Context) error {
	res, err := models.GetAllComments()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func StoreComment(c echo.Context) error {
	comment := c.FormValue("comment")
	userId := c.FormValue("user_id")

	res, err := models.StoreComment(comment, userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteComment(c echo.Context) error {
	id := c.Param("id")

	res, err := models.DeleteComment(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}