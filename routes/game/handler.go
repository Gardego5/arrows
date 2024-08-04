package game

import (
	"fmt"

	"github.com/Gardego5/arrows/lib"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type hWS struct {
	*lib.State
	*websocket.Upgrader
}

func (h hWS) handle(c echo.Context) error {
	ws, err := h.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		err = ws.WriteJSON(map[string]any{"hi": "hello"})
		if err != nil {
			c.Logger().Error(err)
		}

		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
