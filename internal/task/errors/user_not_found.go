package errors

type UserNotFoundError struct {
	// Message is the error message
	Message string
	Err     error
}

func (e UserNotFoundError) Error() string {
	return e.Message
}
