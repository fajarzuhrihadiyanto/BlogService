package _struct

// HTTPFieldError
// This type is used to define response field error message
type HTTPFieldError struct {
	Field   string
	Message string
}

// HTTPBodyResponse
// This type is used to define http body response
type HTTPBodyResponse struct {
	Message string
	Data    interface{}
	Errors  []HTTPFieldError
}
