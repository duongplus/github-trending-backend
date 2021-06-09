package middleware

import (
	"github.com/labstack/echo/v4"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//logic
			return nil
		}
	}
}
