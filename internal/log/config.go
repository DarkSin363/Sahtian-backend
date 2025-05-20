package log

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	FieldKeyModule         = "module"
	FieldKeyFunction       = "function"
	FieldKeyUserID         = "userID"
	FieldKeyRequestPath    = "path"
	FieldKeyRequestMethod  = "method"
	FieldKeyRequestBody    = "request"
	FieldKeyResponseBody   = "response"
	FieldKeyResponseStatus = "status"
	FieldKeySerName        = "srvname"
	FieldKeyDuration       = "duration"
	FieldKeyStatus         = "status"
	FieldKeyClientIP       = "clientIP"
	FieldKeyRequestID      = "requestID"
	FieldKeyErrorStack     = "error_stack"
)

type Format string

const (
	TextFormat Format = "text"
	JsonFormat Format = "json"
)

const (
	outputStderr = "stderr"
	outputStdout = "stdout"
	outputNull   = "null"
)

const (
	defaultFileMaxSize    = 100
	defaultFileMaxAge     = 28
	defaultFileMaxBackups = 3
)

type Config struct {
	Level  string `mapstructure:"level"`
	Format Format `mapstructure:"format"`
	Output string `mapstructure:"output"`
	Name   string `mapstructure:"name"`
}

func init() { //nolint:gochecknoinits
	cfg := &Config{
		Level:  "info",
		Format: TextFormat,
		Output: outputStdout,
		Name:   "",
	}

	_ = Init(cfg)
}

func Init(cfg *Config) error {
	level, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		return errors.Wrapf(err, "parse log level")
	}

	loggerCfg := &zap.Config{
		Level:            level,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	plain, err := loggerCfg.Build(zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		plain = zap.NewNop()
	}
	defaultLogger = plain

	return nil
}

func newJSONEncoder() zapcore.Encoder {
	encConf := zapcore.EncoderConfig{
		MessageKey:          "message",
		LevelKey:            "severity",
		TimeKey:             "@timestamp",
		NameKey:             "facility",
		CallerKey:           "caller",
		FunctionKey:         zapcore.OmitKey,
		StacktraceKey:       "stacktrace",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.LowercaseLevelEncoder,
		EncodeTime:          zapcore.RFC3339TimeEncoder,
		EncodeDuration:      zapcore.MillisDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          zapcore.FullNameEncoder,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "",
	}

	jsonEncoder := zapcore.NewJSONEncoder(encConf)

	return jsonEncoder
}

func newTextEncoder() zapcore.Encoder {
	encConf := zapcore.EncoderConfig{
		MessageKey:          "M",
		LevelKey:            "L",
		TimeKey:             "T",
		NameKey:             "N",
		CallerKey:           "C",
		FunctionKey:         zapcore.OmitKey,
		StacktraceKey:       "S",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.CapitalColorLevelEncoder,
		EncodeTime:          zapcore.RFC3339TimeEncoder,
		EncodeDuration:      zapcore.StringDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          zapcore.FullNameEncoder,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "\t",
	}

	jsonEncoder := zapcore.NewConsoleEncoder(encConf)

	return jsonEncoder
}

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    encodeLevel(),
	EncodeTime:     zapcore.RFC3339TimeEncoder,
	EncodeDuration: zapcore.MillisDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
}

type nullOutput struct {
}

func (n2 *nullOutput) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (n2 *nullOutput) Sync() error {
	return nil
}
