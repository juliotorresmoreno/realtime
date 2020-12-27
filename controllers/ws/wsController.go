package ws

import (
	WS "github.com/juliotorresmoreno/realtime/services/ws"
	"github.com/labstack/echo"
)

// ConfigureWSRouter .
func ConfigureWSRouter(g *echo.Group) {
	g.GET("", GETIndex)
}

func GETIndex(c echo.Context) error {
	WS.Upgrade(c)
	return nil
}

func POSTIndex(c echo.Context) error {
	WS.SendMessage(&WS.Message{
		UserID:  "",
		Content: map[string]string{},
	})
	return nil
}
