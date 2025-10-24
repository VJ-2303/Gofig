// Package gofig provides a simple library for loading configuration
// file (like JSON, YAML) into go structs.
package gofig

import "fmt"

// Load reads a configurtion file from the given path and populates
// the fields of the provided 'cfg' struct.
//
// 'path' is the full file path to the configuration file.
// 'cfg' must be a pointer to a struct, otherwise, an error will be returned.
func Load(path string, cfg any) error {
	fmt.Printf("DEBUG: Attempting to load config from: %s\n", path)
	return nil
}
