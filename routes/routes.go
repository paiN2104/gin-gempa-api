package routes

import (
	"net/http"
	"api_gempa/controllers"
	// "api_gempa/middleware"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Guys!")
	})
	e.POST("/login", controllers.CheckLogin)
	e.GET("/login/:id", controllers.FetchLoginById)
	e.PATCH("/update", controllers.UpdateUser)
	e.POST("/register", controllers.StoreUser)
	e.POST("/comment", controllers.StoreComment)
	e.GET("/comment", controllers.FetchAllComment)
	e.DELETE("/comment/:id", controllers.DeleteComment)
	e.GET("/gempa", controllers.FetchGempaDirasakanApi)
	e.GET("/gempaterkini", controllers.FetchGempaTerkiniApi)
	// e.GET("/hash/:password", controllers.GenerateHashPassword)

	// Routes with middleware
	// r := e.Group("/api")
	// r.Use(middleware.JWT([]byte("secretsauce")))
	// r.GET("/gempa", controllers.GetGempa)

	return e
}
