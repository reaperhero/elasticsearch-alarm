package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/reaperhero/elasticsearch-alarm/pkg/dto"
	"github.com/reaperhero/elasticsearch-alarm/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

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
	ins := h.service.ListAlarmInstance(req.Page, req.Size)
	return e.JSON(200, handleErr(nil, ins))
}

func (h *httphandler) createInstance(e echo.Context) error {
	var (
		req dto.DtoAlarmInstance
		err error
	)

	_ = e.Bind(&req)
	if err = validator.New().Struct(req); err != nil {
		logrus.WithField("validator", err).Error()
		return e.JSON(200, handleErr(errors.ErrRequestParam, nil))
	}
	err = h.service.CreateAlarmInstance(req)
	return e.JSON(200, handleErr(err, nil))
}

func (h *httphandler) deleteInstance(e echo.Context) error {
	id := e.Param("id")
	ident, err := strconv.Atoi(id)
	err = h.service.DeleteAlarmConfig(ident)
	return e.JSON(200, handleErr(err, nil))
}

func (h *httphandler) updateInstance(e echo.Context) error {
	var (
		req dto.DtoAlarmInstance
		err error
	)
	id := e.Param("id")
	ident, err := strconv.Atoi(id)
	_ = e.Bind(&req)
	err = h.service.UpdateAlarmInstance(ident, req)
	return e.JSON(200, handleErr(err, nil))
}
