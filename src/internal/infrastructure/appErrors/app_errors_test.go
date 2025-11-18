package appErrors

import (
	"errors"
	"fmt"
	"testing"
)


func TestAppError_Error(t *testing.T) {
	err := New(EnvError, "переменная не установлена")
	expected := "[env_error] переменная не установлена"
	if err.Error() != expected {
		t.Errorf("Expected %q, got %q", expected, err.Error())
	}
}

func TestAppError_ErrorWithWrapped(t *testing.T) {
	originalErr := fmt.Errorf("внутренняя ошибка")
	err := Wrap(EnvError, "обертка", originalErr)

	expected := "[env_error] обертка: внутренняя ошибка"
	if err.Error() != expected {
		t.Errorf("Expected %q, got %q", expected, err.Error())
	}
}

func TestAppError_Unwrap(t *testing.T) {
	originalErr := fmt.Errorf("внутренняя ошибка")
	err := Wrap(EnvError, "обертка", originalErr)

	unwrapped := err.Unwrap()
	if unwrapped != originalErr {
		t.Errorf("Expected unwrapped error to be %v, got %v", originalErr, unwrapped)
	}
}

func TestAppError_WithErrorsIs(t *testing.T) {
	originalErr := fmt.Errorf("внутренняя ошибка")
	err := Wrap(EnvError, "обертка", originalErr)

	if !errors.Is(err, originalErr) {
		t.Errorf("Expected errors.Is to return true")
	}
}

func TestAppError_WithErrorsAs(t *testing.T) {
	err := New(EnvError, "переменная не установлена")

	var appErr *AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("Expected errors.As to return true")
	}

	if appErr.Type != EnvError {
		t.Errorf("Expected Type to be %s, got %s", EnvError, appErr.Type)
	}

	if appErr.Message != "переменная не установлена" {
		t.Errorf("Expected Message to be 'переменная не установлена', got %s", appErr.Message)
	}
}

func TestAppError_WithNestedWrapping(t *testing.T) {
	originalErr := fmt.Errorf("оригинальная ошибка")
	wrapped1 := Wrap(HttpError, "обертка 1", originalErr)
	wrapped2 := Wrap(EnvError, "обертка 2", wrapped1)

	var appErr *AppError
	if !errors.As(wrapped1, &appErr) {
		t.Fatalf("Expected errors.As to return true")
	}

	if appErr.Type != HttpError {
		t.Errorf("Expected inner error type to be %s, got %s", HttpError, appErr.Type)
	}

	if !errors.Is(wrapped2, originalErr) {
		t.Errorf("Expected errors.Is to return true for original error")
	}
}
