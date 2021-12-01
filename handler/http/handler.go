package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/reaperhero/elasticsearch-alarm/pkg/dto"
	"github.com/reaperhero/elasticsearch-alarm/pkg/errors"
	"github.com/reaperhero/elasticsearch-alarm/pkg/service"
	"github.com/sirupsen/logrus"
)

type httphandler struct {
	service service.WebService
}

func NewHttpHandler() httphandler {
	return httphandler{
		service: service.NewWebService(),
	}
}

func (h *httphandler) queryInstance(e echo.Context) error {
	var (
		req dto.PageSize
		err error
	)
	_ = e.Bind(&req)
	if err = validator.New().Struct(req); err != nil {
		logrus.WithField("validator", err).Info()
		return e.JSON(200, handleErr(errors.ErrRequestParam, nil))
	}
	if req.Size == 0 || req.Page == 0 {
		req.Size, req.Page = 15, 0
	}
	configs, err := h.service.GetAlarmConfig(req.Page, req.Size)
	return e.JSON(200, handleErr(err, configs))
}

func (h *httphandler) createAlarmConfig(e echo.Context) error {
	var (
		req dto.DtoAlarmConfig
		err error
	)

	_ = e.Bind(req)
	if err = validator.New().Struct(req); err != nil {
		return e.JSON(200, handleErr(errors.ErrRequestParam, nil))
	}
	err = h.service.CreateAlarmConfig(req)
	return e.JSON(200, handleErr(err, nil))
}
