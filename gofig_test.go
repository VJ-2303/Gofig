package gofig

import "testing"

func TestLoad(t *testing.T) {
	type AppConfig struct {
		Port int
		Host string
	}
	var cfg AppConfig

	err := Load("fake-path.json", &cfg)
	if err != nil {
		t.Errorf("Load() returned an unexpected error: %v", err)
	}
}
