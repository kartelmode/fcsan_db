package apperrors

// Error meant to signal unauthorized behaviour.
type ErrPermissionDenied struct {
	Message string
}

func (e *ErrPermissionDenied) Error() string {
	return e.Message
}

// Error meant to signal invalid input.
type ErrInvalidArgument struct {
	Message string
}

func (e *ErrInvalidArgument) Error() string {
	return e.Message
}

// Error meant to signal missing content.
type ErrNotFound struct {
	Message string
}

func (e *ErrNotFound) Error() string {
	return e.Message
}

// Error meant to signal internal failure.
type ErrInternal struct {
	Message string
}

func (e *ErrInternal) Error() string {
	return e.Message
}

// Error meant to signal duplicate entry.
type ErrAlreadyExists struct {
	Message string
}

func (e *ErrAlreadyExists) Error() string {
	return e.Message
}

// Error meant to signal unauthenticated request.
type ErrUnauthenticated struct {
	Message string
}

func (e *ErrUnauthenticated) Error() string {
	return e.Message
}
