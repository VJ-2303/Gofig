package gofig_test

import (
	"strings"
	"testing"

	gofig "github.com/VJ-2303/Gofig"
)

// TestLoad is the test function for our Load function.
func TestLoad_JSON(t *testing.T) {
	type AppConfig struct {
		Port int    `json:"port"`
		Host string `json:"host"`
	}
	var cfg AppConfig

	filepath := "testdata/config.json"
	err := gofig.Load(filepath, &cfg)
	if err != nil {
		t.Fatalf("Load() returned an unexpected error: %v", err)
	}

	if cfg.Port != 8080 {
		t.Errorf("expected Port to be 8080, but go %d", cfg.Port)
	}
	if cfg.Host != "localhost" {
		t.Errorf("expected Host to be 'localhost', but go %s", cfg.Host)
	}
}

// TestLoad_UnSupportedFile tests that Load returns an error
// for an unsupported file extension
func TestLoad_UnSupportedFile(t *testing.T) {
	var cfg struct{}

	filepath := "testdata/config.txt"
	err := gofig.Load(filepath, &cfg)

	if err == nil {
		t.Fatal("Expected unsupported filetype Error, but got nil")
	}
	expectedErrMsg := "unsupported config file type: .txt"
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Errorf("expected error to contain: %q, but got %q", expectedErrMsg, err.Error())
	}
}
