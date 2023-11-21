package helper

type BadRequestError struct {
	err error
}

func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{
		err: err,
	}
}

func (e *BadRequestError) Error() string {
	return e.err.Error()
}
