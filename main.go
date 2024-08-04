//go:build !wasm
package main

import (
	"embed"
	"fmt"

	"github.com/Gardego5/arrows/lib"
	"github.com/Gardego5/arrows/routes/game"
	"github.com/Gardego5/arrows/routes/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static/*
var static embed.FS

func main() {
	e := echo.New()

	state, err := lib.LoadState()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.Gzip(),
		middleware.CORS(),
	)

	game.Routes(state, e.Group("/game"))
	pages.Routes(state, e.Group(""))

	e.StaticFS("/static", echo.MustSubFS(static, "static"))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", state.Environment.Port)))
}
