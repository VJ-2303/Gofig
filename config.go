// Package Gofig provides utilities for loading configuration files.
package gofig

import (
	"encoding/json"
	"os"
)

// Load reads a configuration file from the given path and populates
// the fields of the provided 'cfg' struct.
//
// 'filepath' is the path to the configuration file
// 'cfg' must be a pointer to a struct
func Load(filepath string, cfg interface{}) error {
	// Open the configuration file
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	// Ensure the file is closed when the function exists
	defer file.Close()

	// Create a new JSON decodeer that reads from our file
	decoder := json.NewDecoder(file)

	// Decode the JSON into the provided 'cfg' struct
	// 'cfg' must be an Pointer !
	if err := decoder.Decode(cfg); err != nil {
		return err
	}
	return nil
}
