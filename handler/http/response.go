package http

import (
	"github.com/reaperhero/elasticsearch-alarm/pkg/errors"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func handleErr(err error, data interface{}) response {
	resp := response{
		Code: 1000,
		Msg:  "ok",
		Data: nil,
	}
	switch err {
	case errors.ErrRequestParam:
		resp.Code = resp.Code + 1
		resp.Msg = err.Error()
	case errors.ErrDbOperation:
		resp.Code = resp.Code + 2
		resp.Msg = err.Error()
	case errors.ErrDbRecord:
		resp.Code = resp.Code + 3
		resp.Msg = err.Error()
	case errors.ErrDbExistRecord:
		resp.Code = resp.Code + 4
		resp.Msg = err.Error()
	default:
		resp.Data = data
	}
	return resp
}
