package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reaperhero/elasticsearch-alarm/pkg/service"
)

type httphandler struct {
	service service.WebService
}

func NewHttpHandler(webService service.WebService) httphandler {
	return httphandler{
		service: webService,
	}
}

func RunHttpserver(webService service.WebService) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	handler := NewHttpHandler(webService)
	configRouter(e.Group("/alarm/config"), handler)
	instanceRouter(e.Group("/alarm/instance"), handler)
	e.Logger.Fatal(e.Start(":80"))
}

func configRouter(e *echo.Group, handler httphandler) {
	e.POST("/create", handler.createAlarmConfig)
	e.GET("/list", handler.queryAlarmConfig)
	e.DELETE("/delete/:id", handler.deleteAlarmConfig)
	e.PUT("/update/:id", handler.updateAlarmConfig)
}

func instanceRouter(e *echo.Group, handler httphandler) {
	e.POST("/create", handler.createInstance)
	e.GET("/list", handler.queryInstance)
	e.DELETE("/delete/:id", handler.deleteInstance)
	e.PUT("/update/:id", handler.updateInstance)
}
