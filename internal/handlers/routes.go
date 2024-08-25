package handlers

import "github.com/labstack/echo/v4"

func InitHandlers(e *echo.Echo) {
	e.Static("/static", "static")
	e.File("/favicon.ico", "static/imgs/favicon.ico")
	e.File("/robots.txt", "static/txt/robots.txt")
	e.File("/sitemap.xml", "static/txt/sitemap.xml")
	e.GET("/", Home)
	e.GET("/services", Services)
	e.GET("/contact", Contact)
	e.POST("/contact", Contact)
}
