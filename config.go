// Package Gofig provides utilities for loading configuration files.
package gofig

import "fmt"

// Load reads a configuration file from the given path and populates
// the fields of the provided 'cfg' struct.
//
// 'filepath' is the path to the configuration file
// 'cfg' must be a pointer to a struct
func Load(filepath string, cfg interface{}) error {
	fmt.Println("Loading config from:", filepath)
	return nil
}
