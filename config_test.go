package gofig_test

import (
	"testing"

	gofig "github.com/VJ-2303/Gofig"
)

// TestLoad is the test function for our Load function.
func TestLoad(t *testing.T) {
	type AppConfig struct {
		Port int
		Host string
	}
	var cfg AppConfig

	err := gofig.Load("dummy/path.json", &cfg)
	if err != nil {
		t.Errorf("Load() returned an unexpected error: %v", err)
	}
}
