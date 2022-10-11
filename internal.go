package errors

import (
	"google.golang.org/grpc/codes"
)

type internalError struct {
	err         error
	description string
}

func (e *internalError) RPCCode() codes.Code {
	return codes.Internal
}

func (e *internalError) Error() string {
	return e.err.Error()
}

func (e *internalError) Description() string {
	return e.description
}

func Internal(err error, format string, variables ...interface{}) *internalError {
	return &internalError{
		err:         err,
		description: formatter(format, variables...),
	}
}
