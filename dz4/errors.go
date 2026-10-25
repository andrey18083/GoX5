package main

import (
	"errors"
	"fmt"
)

var ErrNetworkTimeout = errors.New("сетевой таймаут")

type CustomError struct {
	Channel string
	Code    int
	Err     error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("ошибка в канале [%s] (код %d): %v", e.Channel, e.Code, e.Err)
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

func (e *CustomError) Is(target error) bool {
	t, ok := target.(*CustomError)
	if !ok {
		return false
	}
	return e.Channel == t.Channel && e.Code == t.Code
}
