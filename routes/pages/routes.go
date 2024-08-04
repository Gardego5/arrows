package pages

import (
	"github.com/Gardego5/arrows/lib"
	"github.com/labstack/echo/v4"
)

func Routes(s *lib.State, e *echo.Group) {
	e.GET("/", hIndex{s}.handle)
}
