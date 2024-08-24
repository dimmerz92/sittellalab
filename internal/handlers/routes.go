package handlers

import "github.com/labstack/echo/v4"

func InitHandlers(e *echo.Echo) {
	e.Static("/static", "static")
	e.GET("/", Home)
	e.GET("/services", Services)
	e.GET("/contact", Contact)
	e.POST("/contact", Contact)
}
