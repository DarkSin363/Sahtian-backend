package errors

import (
	"github.com/BigDwarf/sahtian/internal/log"
	"runtime/debug"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Recover(outErr *error, module string) {
	if r := recover(); r != nil {
		if outErr == nil {
			outErr = new(error)
		}

		switch x := r.(type) {
		case string:
			*outErr = errors.Errorf("%s", x)
		case error:
			*outErr = x
		default:
			*outErr = errors.Errorf("%+v", x)
		}

		stack := string(debug.Stack())
		log.With(
			zap.String(log.FieldKeyModule, module),
			zap.String("debug_stack", stack),
			zap.Error(*outErr),
		).Sugar().Errorf("recover after PANIC")
	}
}
