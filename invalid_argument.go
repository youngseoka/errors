package errors

import (
	"google.golang.org/grpc/codes"
)

type invalidArgument struct {
	err         error
	description string
}

func (e *invalidArgument) RPCCode() codes.Code {
	return codes.InvalidArgument
}

func (e *invalidArgument) Error() string {
	return e.err.Error()
}

func (e *invalidArgument) Description() string {
	return e.description
}

func InvalidArgument(err error,format string, variables ...interface{}) *invalidArgument {
	return &invalidArgument{
		err:         err,
		description: formatter(format, variables...),
	}
}
