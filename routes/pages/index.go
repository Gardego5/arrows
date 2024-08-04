package pages

import (
	"net/http"

	"github.com/Gardego5/arrows/lib"
	"github.com/Gardego5/arrows/lib/components"
	. "github.com/Gardego5/htmdsl"
	"github.com/labstack/echo/v4"
)

type hIndex struct{ *lib.State }

func (h hIndex) handle(c echo.Context) error {
	return c.Stream(http.StatusOK, "text/html", components.Layout("Home", nil,
		H1{lib.CN("text-blue-500"), "Arrows!"},
		Iframe{Attrs{{"src", "/game/app"}, {"width", "720"}, {"height", "480"}}},
	).Reader())
}
