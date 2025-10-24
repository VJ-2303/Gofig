// Package Gofig provides utilities for loading configuration files.
package gofig

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
		// Returns an clear error message when the file type is unsupported
		return fmt.Errorf("unsupported config file type: %s", ext)
	}
}

// parseJSON is a private helper function to decode JSON from a file
func parseJSON(file *os.File, cfg interface{}) error {
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}
	return nil
}
