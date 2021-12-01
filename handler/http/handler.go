package http

import "github.com/labstack/echo/v4"

type httphandler struct {
}

func NewHttpHandler() httphandler {
	return httphandler{}
}

func (h *httphandler) getInstance(e echo.Context) error {
	return e.JSON(200, "ok")
}
