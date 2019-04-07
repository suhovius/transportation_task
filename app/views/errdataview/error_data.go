package errdataview

// ErrorData contains error information for API response
type ErrorData struct {
	Message string `json:"error_message"`
}

// New returns new ErrorData view instance
func New(Message string) ErrorData {
	return ErrorData{Message: Message}
}
