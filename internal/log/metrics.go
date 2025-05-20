package log

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	subsystem string = "logger"

	labelLevelDebug = "debug"
	labelLevelInfo  = "info"
	labelLevelError = "error"
	labelLevelWarn  = "warn"
	labelLevelFatal = "fatal"
	labelLevelPanic = "panic"
)

func UseMetrics(namespace string) {
	metrics := newMetrics(namespace)

	core := zapcore.RegisterHooks(defaultLogger.Core(), metrics.hook)

	defaultLogger = zap.New(core)
}

type metrics struct {
	counter *prometheus.CounterVec
}

func newMetrics(namespace string) *metrics {
	m := &metrics{}
	m.counter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "message_total",
			Help:      "Total number of log messages.",
		},
		[]string{"level"},
	)

	return m
}

func (m *metrics) hook(entry zapcore.Entry) error {
	switch entry.Level {
	case zapcore.DebugLevel:
		m.counter.WithLabelValues(labelLevelDebug).Inc()
	case zapcore.InfoLevel:
		m.counter.WithLabelValues(labelLevelInfo).Inc()
	case zapcore.ErrorLevel:
		m.counter.WithLabelValues(labelLevelError).Inc()
	case zapcore.WarnLevel:
		m.counter.WithLabelValues(labelLevelWarn).Inc()
	case zapcore.PanicLevel, zapcore.DPanicLevel:
		m.counter.WithLabelValues(labelLevelPanic).Inc()
	case zapcore.FatalLevel:
		m.counter.WithLabelValues(labelLevelFatal).Inc()
	}

	return nil
}
