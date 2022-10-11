package errors

import "google.golang.org/grpc/codes"

type Error interface {
	error
	Description() string
	RPCCode() codes.Code
}

var _ Error = &forbiddenError{}
var _ Error = &notFoundError{}
var _ Error = &internalError{}
var _ Error = &invalidArgument{}
var _ Error = &unauthorizedError{}
var _ Error = &conflictError{}
var _ Error = &unprocessableError{}

func FromError(err error) Error {
	e, ok := err.(Error)
	if !ok {
		return Internal(err, "unknown internal error")
	}

	return e
}