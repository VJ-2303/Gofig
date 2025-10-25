package gofig

import (
	"errors"
	"testing"
)

func TestLoad(t *testing.T) {
	type AppConfig struct {
		Port int
		Host string
	}
	var cfg AppConfig

	// Test 1: Success Path
	t.Run("success: valid pointer and file exists", func(t *testing.T) {
		err := Load("testdata/dummy.json", &cfg)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})

	// Test 2: 'cfg' is not a pointer
	t.Run("fail: 'cfg' is not a pointer", func(t *testing.T) {
		err := Load("testdata/dummy.txt", cfg)
		if err != ErrNotAPointer {
			t.Errorf("Expected error %v, but got: %v", ErrNotAPointer, err)
		}
	})

	// Test 3: File does not exists
	t.Run("fail:file does not exists", func(t *testing.T) {
		err := Load("non-exists-file.json", &cfg)
		if err != ErrNotFound {
			t.Errorf("Expected error %v, but got: %v", ErrNotFound, err)
		}
	})

	// Test 4: Unsupported file format
	t.Run("fail: unsupported file format", func(t *testing.T) {
		err := Load("testdata/dummy.txt", &cfg)
		if !errors.Is(err, ErrUnsupportedFormat) {
			t.Errorf("Expected error %v, but got: %v", ErrUnsupportedFormat, err)
		}
	})
}
