package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"sittellalab.com.au/internal/email"
	"sittellalab.com.au/templates"
)

func Contact(ctx echo.Context) error {
	switch ctx.Request().Method {

	case http.MethodPost:
		// parse mail form
		if err := email.SendMail(ctx); err != nil {
			return Render(ctx, http.StatusOK, templates.Contact(templates.ContactProps{Error: err.Error()}))
		}
		return Render(ctx, http.StatusOK, templates.Contact(templates.ContactProps{Success: true}))

	default:
		return Render(ctx, http.StatusOK, templates.Contact(templates.ContactProps{}))
	}
}
