package error

type UnauthorizedError struct {
	err error
}

func NewUnauthorizedError(err error) *UnauthorizedError {
	return &UnauthorizedError{
		err: err,
	}
}

func (e *UnauthorizedError) Error() string {
	return e.err.Error()
}
