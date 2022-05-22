package storage

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrNotFoundId = errors.New("not found record with that id")
	// ErrNotFoundUDelete is the error when the record with that id is not found to delete
	ErrNotFoundDelete = errors.New("not found record to delete")
	// ErrNotFoundUpdate is the error when the record with that id is not found to update
	ErrNotFoundUpdate = errors.New("not found record to update")
)
