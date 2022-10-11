package errors

import "google.golang.org/grpc/codes"

type unprocessableError struct {
	err         error
	description string
}

func (e *unprocessableError) RPCCode() codes.Code {
	return codes.FailedPrecondition
}

func (e *unprocessableError) Error() string {
	return e.err.Error()
}

func (e *unprocessableError) Description() string {
	return e.description
}

func Unprocessable(err error, format string, variables ...interface{}) *unprocessableError {
	return &unprocessableError{
		err:         err,
		description: formatter(format, variables...),
	}
}
