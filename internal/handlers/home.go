package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"sittellalab.com.au/templates"
)

func Home(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, templates.Home())
}
