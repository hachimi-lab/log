package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Builder        = zap.Option
	Core           = zapcore.Core
	Clock          = zapcore.Clock
	Entry          = zapcore.Entry
	CheckWriteHook = zapcore.CheckWriteHook
	CheckedEntry   = zapcore.CheckedEntry
	LevelEnabler   = zapcore.LevelEnabler
)

func BuildWithWrapCore(fn func(Core) Core) Builder {
	return zap.WrapCore(fn)
}

func BuildWithHooks(hooks ...func(entry Entry) error) Builder {
	return zap.Hooks(hooks...)
}

func BuildWithFields(fs ...Field) Builder {
	return zap.Fields(fs...)
}

func BuildWithErrorOutput(w WriteSyncer) Builder {
	return zap.ErrorOutput(w)
}

func BuildWithDevelopment() Builder {
	return zap.Development()
}

func BuildWithAddCaller() Builder {
	return zap.AddCaller()
}

func BuildWithCaller(enabled bool) Builder {
	return zap.WithCaller(enabled)
}

func BuildWithAddCallerSkip(skip int) Builder {
	return zap.AddCallerSkip(skip)
}

func BuildWithAddStacktrace(lvl LevelEnabler) Builder {
	return zap.AddStacktrace(lvl)
}

func BuildWithIncreaseLevel(lvl Level) Builder {
	return zap.IncreaseLevel(lvl)
}

func BuildWithFatalHook(hook CheckWriteHook) Builder {
	return zap.WithFatalHook(hook)
}

func BuildWithClock(clock Clock) Builder {
	return zap.WithClock(clock)
}
