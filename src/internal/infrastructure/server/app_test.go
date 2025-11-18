package server

import (
	"golang-web-server/src/internal/infrastructure/config"
	"testing"
)


func TestNewApp(t *testing.T) {
	cfg := &config.Config{
		Server: config.Server{
			Host:        "localhost",
			Port:        "8080",
			Environment: "development",
		},
	}

	app := NewApp(cfg)

	if app == nil {
		t.Fatalf("Expected app to be created, got nil")
	}
}

