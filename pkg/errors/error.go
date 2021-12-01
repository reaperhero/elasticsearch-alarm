package errors

import "errors"

// web
var (
	ErrRequestParam = errors.New("request param error")
	ErrDbOperation  = errors.New("db error")
)
