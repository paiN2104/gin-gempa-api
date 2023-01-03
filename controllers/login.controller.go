package controllers

import (
	"net/http"
	// "time"
	"api_gempa/helpers"
	"api_gempa/models"

	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

//get login data by id
func FetchLoginById(c echo.Context) error {

	id := c.Param("id")

	result, err := models.FetchUserById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

//get login data
func FetchLogin(c echo.Context) error {
	res, err := models.FetchLogin()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}


func CheckLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := models.CheckLogin(email, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	// if !res{
	// 	return echo.ErrUnauthorized
	// }

	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["level"] = "application"
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// mytoken, err := token.SignedString([]byte("secretsauce"))
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	// }

	return c.JSON(http.StatusOK,
		map[string]interface{}{
			"id": res,
			"message": "Login successful",
		})
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func StoreUser(c echo.Context) error {
	name := c.FormValue("name")
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	status := c.FormValue("status")
	image := c.FormValue("image")

	res, err := models.Register(name, username, password, email, status, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res{
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK,
		map[string]string{
			"message": "Register successful",
		})
}

//update user by id
func UpdateUser(c echo.Context) error {
	id := helpers.ConvertStringToInt(c.FormValue("id"))
	name := c.FormValue("name")
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	status := c.FormValue("status")
	image := c.FormValue("image")

	var pwd = password

	pass, err := helpers.HashPassword(pwd)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := models.UpdateUser(id, name, username, pass, email, status, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res{
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK,
		map[string]string{
			"message": "Update successful",
		})
}
