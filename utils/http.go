package utils

import (
	"errors"

	"github.com/labstack/echo"
)

func PathPrefix(prefix string) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			length := len(prefix)
			if c.Request().URL.Path[:length] == prefix {
				nval := c.Request().URL.Path[length:]
				c.Request().URL.Path = nval
				return h(c)
			}
			return errors.New("PathPrefix, el valor no es valido para " + prefix)
		}
	}
}
