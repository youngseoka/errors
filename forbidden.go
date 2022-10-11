package errors

import (
	"google.golang.org/grpc/codes"
)

type forbiddenError struct {
	err         error
	description string
}

func (e *forbiddenError) RPCCode() codes.Code {
	return codes.PermissionDenied
}

func (e *forbiddenError) Error() string {
	return e.err.Error()
}

func (e *forbiddenError) Description() string {
	return e.description
}

func Forbidden(err error, format string, variables ...interface{}) *forbiddenError {
	return &forbiddenError{
		err:         err,
		description: formatter(format, variables...),
	}
}
