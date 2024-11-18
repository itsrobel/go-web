package handlers

import (
	"web/internal/templates"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return templates.Home().Render(c.Request().Context(), c.Response().Writer)
}

func AboutHandler(c echo.Context) error {
	return templates.About().Render(c.Request().Context(), c.Response().Writer)
}
