package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sentinal Contact Form API Server")
	})
}
