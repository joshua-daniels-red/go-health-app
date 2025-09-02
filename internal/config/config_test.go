package config

import (
	"os"
	"testing"
)

func TestLoad_DefaultPort(t *testing.T) {
	os.Unsetenv("PORT")
	cfg, _ := Load()
	if cfg.Port != "8080" {
		t.Errorf("Expected default port 8080, got %s", cfg.Port)
	}
}

func TestLoad_CustomPort(t *testing.T) {
	os.Setenv("PORT", "9000")
	defer os.Unsetenv("PORT")

	cfg, _ := Load()
	if cfg.Port != "9000" {
		t.Errorf("Expected port 9000, got %s", cfg.Port)
	}
}
