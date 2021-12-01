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
	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

func setRouter(e *echo.Group, handler httphandler) {
	e.GET("/get",handler.getInstance)
}
