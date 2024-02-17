package common

import "errors"

var (
	ErrDataNotFound = errors.New("data nof found")
	ErrLoginFail    = errors.New("username or password invalid")
)
