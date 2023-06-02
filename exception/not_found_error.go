package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(err string) NotFoundError {
	return NotFoundError{err}
}
