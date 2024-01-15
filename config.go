package log

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Config           func(config *zap.Config)
	LevelEncoder     = zapcore.LevelEncoder
	TimeEncoder      = zapcore.TimeEncoder
	DurationEncoder  = zapcore.DurationEncoder
	CallerEncoder    = zapcore.CallerEncoder
	NameEncoder      = zapcore.NameEncoder
	ReflectedEncoder = zapcore.ReflectedEncoder
	SamplingConfig   = zap.SamplingConfig
)

func WithLevel(level Level) Config {
	return func(config *zap.Config) {
		config.Level.SetLevel(level)
	}
}

func WithDevelopment(development bool) Config {
	return func(config *zap.Config) {
		config.Development = development
	}
}

func WithDisableCaller(disableCaller bool) Config {
	return func(config *zap.Config) {
		config.DisableCaller = disableCaller
	}
}

func WithDisableStacktrace(disableStacktrace bool) Config {
	return func(config *zap.Config) {
		config.DisableStacktrace = disableStacktrace
	}
}

func WithSampling(sampling *SamplingConfig) Config {
	return func(config *zap.Config) {
		config.Sampling = sampling
	}
}

func WithEncoding(encoding Encoding) Config {
	return func(config *zap.Config) {
		config.Encoding = encoding
	}
}

func WithEncoderMessageKey(encoderMessageKey string) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.MessageKey = encoderMessageKey
	}
}

func WithEncoderLevelKey(encoderLevelKey string) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.LevelKey = encoderLevelKey
	}
}

func WithEncoderTimeKey(encoderTimeKey string) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.TimeKey = encoderTimeKey
	}
}

func WithEncoderNameKey(encoderNameKey string) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.NameKey = encoderNameKey
	}
}

func WithEncoderCallerKey(encoderCallerKey string) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.CallerKey = encoderCallerKey
	}
}

func WithEncoderFunctionKey(encoderFunctionKey string) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.FunctionKey = encoderFunctionKey
	}
}

func WithEncoderStacktraceKey(encoderStacktraceKey string) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.StacktraceKey = encoderStacktraceKey
	}
}

func WithEncoderLineEnding(encoderLineEnding string) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.LineEnding = encoderLineEnding
	}
}

func WithEncoderLevel(encoderLevel LevelEncoder) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.EncodeLevel = encoderLevel
	}
}

func WithEncoderTime(encoderTime TimeEncoder) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.EncodeTime = encoderTime
	}
}

func WithEncoderDuration(encoderDuration DurationEncoder) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.EncodeDuration = encoderDuration
	}
}

func WithEncoderCaller(encoderCaller CallerEncoder) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.EncodeCaller = encoderCaller
	}
}

func WithEncoderName(encoderName NameEncoder) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.EncodeName = encoderName
	}
}

func WithEncoderNewReflectedEncoder(encoderNewReflectedEncoder func(io.Writer) ReflectedEncoder) Config {
	return func(config *zap.Config) {
		config.EncoderConfig.NewReflectedEncoder = encoderNewReflectedEncoder
	}
}

func WithOutputPaths(outputPaths ...string) Config {
	return func(config *zap.Config) {
		config.OutputPaths = outputPaths
	}
}

func WithErrorOutputPaths(errorOutputPaths ...string) Config {
	return func(config *zap.Config) {
		config.ErrorOutputPaths = errorOutputPaths
	}
}

func WithInitialFields(initialFields map[string]interface{}) Config {
	return func(config *zap.Config) {
		config.InitialFields = initialFields
	}
}
