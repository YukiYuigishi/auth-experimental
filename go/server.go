package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func login(c echo.Context) error {
	return c.String(http.StatusOK, "success")
}
func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, echo.Version)
	})

	e.POST("/login", login)

	e.Logger.Fatal(e.Start(":5000"))

}
