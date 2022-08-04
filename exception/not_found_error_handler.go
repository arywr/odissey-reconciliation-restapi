package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{Error: message}
}
