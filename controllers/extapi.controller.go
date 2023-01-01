package controllers

import (
	"net/http"
	// "api_gempa/models"
	"io/ioutil"
	"github.com/labstack/echo/v4"
)


func FetchGempaTerkiniApi(c echo.Context) error {
    // Make the request
    response, err := http.Get("https://data.bmkg.go.id/DataMKG/TEWS/autogempa.json")
    if err != nil {
        return err
    }
    defer response.Body.Close()

    // Read the response
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return err
    }

    return c.String(http.StatusOK, string(data))
}

func FetchGempaDirasakanApi(c echo.Context) error {
    // Make the request
    response, err := http.Get("https://data.bmkg.go.id/DataMKG/TEWS/gempadirasakan.json")
    if err != nil {
        return err
    }
    defer response.Body.Close()

    // Read the response
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return err
    }

    return c.String(http.StatusOK, string(data))
}