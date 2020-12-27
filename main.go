package main

import (
	"github.com/juliotorresmoreno/realtime/controllers/ws"
	"github.com/juliotorresmoreno/realtime/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	ws.ConfigureWSRouter(e.Group("/ws", utils.PathPrefix("/ws")))
	e.Logger.Fatal(e.Start(":1323"))
}
