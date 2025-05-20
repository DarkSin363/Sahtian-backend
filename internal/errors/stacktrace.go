package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func StackTrace(err error) string {
	var (
		errStack stackTracer
		stack    string
	)

	if errors.As(err, &errStack) {
		stack = fmt.Sprintf("%+v", errStack.StackTrace())
	}

	return stack
}
