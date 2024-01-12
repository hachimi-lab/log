package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Level            = zapcore.Level
	LevelEnablerFunc = zap.LevelEnablerFunc
)

const (
	DebugLevel  Level = zapcore.DebugLevel
	InfoLevel   Level = zapcore.InfoLevel
	WarnLevel   Level = zapcore.WarnLevel
	ErrorLevel  Level = zapcore.ErrorLevel
	DPanicLevel Level = zapcore.DPanicLevel
	PanicLevel  Level = zapcore.PanicLevel
	FatalLevel  Level = zapcore.FatalLevel
)
