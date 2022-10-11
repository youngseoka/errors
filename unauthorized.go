package errors

import (
	"google.golang.org/grpc/codes"
)

type unauthorizedError struct {
	err         error
	description string
}

func (e *unauthorizedError) RPCCode() codes.Code {
	return codes.Unauthenticated
}

func (e *unauthorizedError) Error() string {
	return e.err.Error()
}

func (e *unauthorizedError) Description() string {
	return e.description
}

func Unauthorized(err error, format string, variables ...interface{}) *unauthorizedError {
	return &unauthorizedError{
		err:         err,
		description: formatter(format, variables...),
	}
}
