package rp

import "errors"

var (
	ErrNotValidPath = errors.New("not valid path")
	ErrInvalidFile = errors.New("invalid file")
	ErrInvalidDir = errors.New("invalid dir")
)