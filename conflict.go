package errors

import (
	"google.golang.org/grpc/codes"
)

type conflictError struct {
	err         error
	description string
}

func (e *conflictError) RPCCode() codes.Code {
	return codes.AlreadyExists
}

func (e *conflictError) Error() string {
	return e.err.Error()
}

func (e *conflictError) Description() string {
	return e.description
}

func Conflict(err error, format string, variables ...interface{}) *conflictError {
	return &conflictError{
		err:         err,
		description: formatter(format, variables...),
	}
}
