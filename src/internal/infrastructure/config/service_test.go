package config

import (
	"errors"
	"golang-web-server/src/internal/infrastructure/appErrors"
	"os"
	"testing"
)


func TestConfigService_GetConfig_Success(t *testing.T){ 
	t.Setenv("API_HOST", "localhost")
	t.Setenv("API_PORT", "8080")
	t.Setenv("ENVIRONMENT", "development")

	cs := NewConfigService()
	cfg, err := cs.GetConfig()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cfg.Server.Host != "localhost"{
		t.Errorf("Expected Host to be 'localhost', got '%s'", cfg.Server.Host)
	}
	if cfg.Server.Port != "8080"{
		t.Errorf("Expected Port to be '8080', got '%s'", cfg.Server.Port)
	}
	if cfg.Server.Environment != "development"{
		t.Errorf("Expected Environment to be 'development', got '%s'", cfg.Server.Environment)
	}
}

func TestConfigService_GetConfig_MissingEnv(t *testing.T) {
	t.Setenv("API_HOST", "") 
	os.Unsetenv("API_HOST")
	os.Unsetenv("API_PORT")
	os.Unsetenv("ENVIRONMENT")

	cs := NewConfigService()
	_, err := cs.GetConfig()

	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}

	var appErr *appErrors.AppError
	if !errors.As(err, &appErr) || appErr.Type != appErrors.EnvError {
		t.Fatalf("Expected AppError with type EnvError, got %v", err)
	}
}