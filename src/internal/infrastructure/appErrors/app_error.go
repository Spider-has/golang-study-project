package appErrors

import "fmt"

type ErrorType = string;

const (
	EnvError ErrorType = "env_error"
	HttpError ErrorType = "http_error"
	DBError ErrorType = "db_error"
	unknown ErrorType = "unknow_error"
)

type AppError struct {
	Type ErrorType
	Message string
	Err error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}


func New( errType ErrorType, message string ) *AppError {
	return &AppError{
		Type: errType,
		Message: message,
	}
}

func(e *AppError) Unwrap() error {
	return e.Err
}

func Wrap( errType ErrorType, message string, err error ) *AppError {
	return &AppError{
		Type: errType,
		Message: message,
		Err: err,
	}
}

