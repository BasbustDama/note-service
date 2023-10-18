package entity

import "errors"

type AppError error

var (
	ErrorInternal   AppError = errors.New("internal server error")
	ErrorNotFound   AppError = errors.New("not found error")
	ErrorBadRequest AppError = errors.New("bad request error")
)
