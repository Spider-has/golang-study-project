package appErrors

import "fmt"


type HTTPError struct {
	Status int
	Title string
	Description string
	Err error
}

func (e *HTTPError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[HTTP %d] %s: %s (внутренняя ошибка: %v)", e.Status, e.Title, e.Description, e.Err)
	}
	return fmt.Sprintf("[HTTP %d] %s: %s", e.Status, e.Title, e.Description)
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

func NewHTTPError(statusCode int, title, details string) *HTTPError {
	return &HTTPError{
		Status: statusCode,
		Title:      title,
		Description:    details,
	}
}

func WrapHTTPError(statusCode int, title, details string, err error) *HTTPError {
	return &HTTPError{
		Status: statusCode,
		Title:      title,
		Description:    details,
		Err:        err,
	}
}