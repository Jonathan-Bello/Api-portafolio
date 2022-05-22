package handlers

import (
	"errors"
)

var (
	// ErrIdRequired is the error when the id is required
	ErrIdRequired  = errors.New("id required and must be greater than 0")
	ErrInvalidBody = errors.New("invalid body to parser json")
)
