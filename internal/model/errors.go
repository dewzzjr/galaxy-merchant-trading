package model

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidQuestion = errors.New("invalid question")
	ErrInvalidUnit     = errors.New("invalid unit")
)
