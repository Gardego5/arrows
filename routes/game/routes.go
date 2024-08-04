package game

import (
	"github.com/Gardego5/arrows/lib"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func Routes(s *lib.State, e *echo.Group) {
	e.GET("/ws", hWS{s, &websocket.Upgrader{}}.handle)
	e.GET("/app", hApp{s}.handle)
}
