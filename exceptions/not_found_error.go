package exceptions

type NotFoundError struct {
	Error string
}

func NewNotFoundError(err string) NotFoundError {
	return NotFoundError{Error: err}
}

type ErrorBadRequest struct {
	Error string
}

func NewErrorBadRequest(err string) ErrorBadRequest {
	return ErrorBadRequest{Error: err}
}

type ErrorUnsupported struct {
	Error string
}

func NewErrorUnsupported(err string) ErrorUnsupported {
	return ErrorUnsupported{Error: err}
}
