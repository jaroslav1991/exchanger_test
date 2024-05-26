package service

import "errors"

var (
	errIncorrectAmount      = errors.New("incorrect amount")
	errZeroOrNegativeAmount = errors.New("zero or negative amount")
)
