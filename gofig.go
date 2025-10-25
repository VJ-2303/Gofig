// Package gofig provides a simple library for loading configuration
// file (like JSON, YAML) into go structs.
package gofig

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

// Load reads a configurtion file from the given path and populates
// the fields of the provided 'cfg' struct.
//
// 'path' is the full file path to the configuration file.
// 'cfg' must be a pointer to a struct, otherwise, an error will be returned.
func Load(path string, cfg any) error {
	// Check if 'cfg' is a pointer to a struct
	v := reflect.ValueOf(cfg)

	// If it's not a Ptr, We return the ErrNotAPointer error
	if v.Kind() != reflect.Ptr {
		return ErrNotAPointer
	}

	// Read the file's raw bytes.
	_, err := os.ReadFile(path)
	if err != nil {
		// Check if the error is 'file not found' and return custom error
		if errors.Is(err, os.ErrNotExist) {
			return ErrNotFound
		}
		return err
	}

	// Get the file extension using filepath.Ext function
	ext := filepath.Ext(path)

	// Decide which parser to use based on the extension
	switch ext {
	case ".json":
		fmt.Printf("DEBUG: Detected JSON file\n")
	default:
		// If the extension is not recognized return the custom error
		return ErrUnsupportedFormat
	}
	return nil
}
