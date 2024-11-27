package errors

type PermissionDeniedError struct {
	// Message is the error message
	Message string
	Err     error
}

func (e PermissionDeniedError) Error() string {
	return e.Message
}
