package errors

import "fmt"

func formatter(format string, variables ...interface{}) string {
	if len(variables) > 0 {
		return fmt.Sprintf(format, variables)
	}

	return format
}
