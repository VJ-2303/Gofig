package gofig

import "errors"

var (
	// ErrNotAPointer is returned when the 'cfg' argument for Load function
	// is not a pointer to a struct.
	ErrNotAPointer = errors.New("gofig: 'cfg' must be a pointer to a struct")

	// ErrNotFound is returned when the configuration file
	// at the specified path is does not exists.
	ErrNotFound = errors.New("gofig: configuration file not found")

	// ErrUnsupportedFormat is returned when the file extension
	// is not supported config format
	ErrUnsupportedFormat = errors.New("gofig: unsupported file format")
)
