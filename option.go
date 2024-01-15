package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Option         = zap.Option
	Core           = zapcore.Core
	Clock          = zapcore.Clock
	Entry          = zapcore.Entry
	CheckWriteHook = zapcore.CheckWriteHook
	CheckedEntry   = zapcore.CheckedEntry
	LevelEnabler   = zapcore.LevelEnabler
)

func BuildWithWrapCore(fn func(Core) Core) Option {
	return zap.WrapCore(fn)
}

func BuildWithHooks(hooks ...func(entry Entry) error) Option {
	return zap.Hooks(hooks...)
}

func BuildWithFields(fs ...Field) Option {
	return zap.Fields(fs...)
}

func BuildWithErrorOutput(w WriteSyncer) Option {
	return zap.ErrorOutput(w)
}

func BuildWithDevelopment() Option {
	return zap.Development()
}

func BuildWithAddCaller() Option {
	return zap.AddCaller()
}

func BuildWithCaller(enabled bool) Option {
	return zap.WithCaller(enabled)
}

func BuildWithAddCallerSkip(skip int) Option {
	return zap.AddCallerSkip(skip)
}

func BuildWithAddStacktrace(lvl LevelEnabler) Option {
	return zap.AddStacktrace(lvl)
}

func BuildWithIncreaseLevel(lvl Level) Option {
	return zap.IncreaseLevel(lvl)
}

func BuildWithFatalHook(hook CheckWriteHook) Option {
	return zap.WithFatalHook(hook)
}

func BuildWithClock(clock Clock) Option {
	return zap.WithClock(clock)
}
