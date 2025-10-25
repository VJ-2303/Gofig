// Package gofig provides a simple library for loading configuration
// file (like JSON, YAML) into go structs.
package gofig

import (
	"fmt"
	"os"
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

	// Check if the file exists, os.Stat gets file info and an error
	if _, err := os.Stat(path); err != nil {
		// Check if the error is specifically a "not exists" error
		if os.IsNotExist(err) {
			// If yes return our custom ErrNotFound error
			return ErrNotFound
		}
		// Might be some other error, (e.g permission)
		return err
	}
	fmt.Printf("DEBUG: Attempting to load config from: %s\n", path)
	return nil
}
