package log

import (
	"go.uber.org/zap"
)

type Logger struct {
	internal *zap.Logger
}

func (slf *Logger) Named(name string) *Logger {
	slf.internal = slf.internal.Named(name)
	return slf
}

func (slf *Logger) Debug(msg string, fields ...Field) {
	slf.internal.Debug(msg, fields...)
}

func (slf *Logger) Info(msg string, fields ...Field) {

	slf.internal.Info(msg, fields...)
}

func (slf *Logger) Warn(msg string, fields ...Field) {
	slf.internal.Warn(msg, fields...)
}

func (slf *Logger) Error(msg string, fields ...Field) {
	slf.internal.Error(msg, fields...)
}

func (slf *Logger) DPanic(msg string, fields ...Field) {
	slf.internal.DPanic(msg, fields...)
}

func (slf *Logger) Panic(msg string, fields ...Field) {
	slf.internal.Panic(msg, fields...)
}

func (slf *Logger) Fatal(msg string, fields ...Field) {
	slf.internal.Fatal(msg, fields...)
}
