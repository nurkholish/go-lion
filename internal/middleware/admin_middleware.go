package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminMiddleware struct{}

func NewAdminMiddleware() *AdminMiddleware {
	return &AdminMiddleware{}
}

func (m *AdminMiddleware) Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role").(string)
		if role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Access denied: Admin role required"})
		}

		return next(c)
	}
}
