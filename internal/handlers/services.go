package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"sittellalab.com.au/templates"
)

func Services(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, templates.Services())
}
