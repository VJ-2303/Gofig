package gofig

import "testing"

func TestLoad(t *testing.T) {
	type AppConfig struct {
		Port int
		Host string
	}
	var cfg AppConfig

	t.Run("success: valid pointer and file exists", func(t *testing.T) {
		err := Load("testdata/dummy.txt", &cfg)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})
	t.Run("fail: 'cfg' is not a pointer", func(t *testing.T) {
		err := Load("testdata/dummy.txt", cfg)
		if err != ErrNotAPointer {
			t.Errorf("Expected error %v, but got: %v", ErrNotAPointer, err)
		}
	})
	t.Run("fail:file does not exists", func(t *testing.T) {
		err := Load("non-exists-file.json", &cfg)
		if err != ErrNotFound {
			t.Errorf("Expected error %v, but got: %v", ErrNotFound, err)
		}
	})
}
