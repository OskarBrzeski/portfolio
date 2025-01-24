package main

import (
	"ob/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Renderer = utils.NewTemplates()
	e.Use(middleware.Logger())

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/test", func(c echo.Context) error {
		return c.Render(200, "success", nil)
	})

	e.GET("/education", func(c echo.Context) error {
		return c.Render(200, "education", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
