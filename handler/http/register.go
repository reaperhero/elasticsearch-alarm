package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunHttpserver() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	group := e.Group("/alarm")
	handler := NewHttpHandler()
	setRouter(group, handler)
	e.Logger.Fatal(e.Start(":80"))
}

func setRouter(e *echo.Group, handler httphandler) {
	e.POST("/instance/create", handler.createAlarmConfig)
	e.GET("/instance/list", handler.queryInstance)
	e.DELETE("/instance/delete/:id", handler.deleteAlarmConfig)
	e.POST("/instance/update/:id", handler.updateAlarmConfig)
}
