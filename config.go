// Package Gofig provides utilities for loading configuration files.
package gofig

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	// ErrUnsupportedFileType is returned when the config file
	// extension is not supported by the loader
	ErrUnsupportedFileType = errors.New("gofig: unsupported file type")

	// ErrParseError is returned when the config file is syntactically invalid
	ErrParseError = errors.New("gofig: parse error")
)

// Load reads a configuration file from the given path and populates
// the fields of the provided 'cfg' struct.
//
// 'filepath' is the path to the configuration file
// 'cfg' must be a pointer to a struct
func Load(filePath string, cfg interface{}) error {
	// Open the configuration file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	// Ensure the file is closed when the function exists
	defer file.Close()

	// Get the file extension
	ext := filepath.Ext(filePath)

	// Use a 'switch' to decide which parser to use
	switch ext {
	case ".json":
		// Delegate to our new JSON parser function
		return parseJSON(file, cfg)
	default:
		return fmt.Errorf("%w: %s", ErrUnsupportedFileType, ext)
	}
}

// parseJSON is a private helper function to decode JSON from a file
func parseJSON(file *os.File, cfg interface{}) error {
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return fmt.Errorf("%w: %v", ErrParseError, err)
	}
	return nil
}
