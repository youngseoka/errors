package errors

import (
	"google.golang.org/grpc/codes"
)

type notFoundError struct {
	err         error
	description string
}

func (e *notFoundError) RPCCode() codes.Code {
	return codes.NotFound
}

func (e *notFoundError) Error() string {
	return e.err.Error()
}

func (e *notFoundError) Description() string {
	return e.description
}

func NotFound(err error, format string, variables ...interface{}) *notFoundError {
	return &notFoundError{
		err:         err,
		description: formatter(format, variables...),
	}
}
