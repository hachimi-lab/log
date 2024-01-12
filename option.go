package log

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Option           func(config *zap.Config)
	LevelEncoder     = zapcore.LevelEncoder
	TimeEncoder      = zapcore.TimeEncoder
	DurationEncoder  = zapcore.DurationEncoder
	CallerEncoder    = zapcore.CallerEncoder
	NameEncoder      = zapcore.NameEncoder
	ReflectedEncoder = zapcore.ReflectedEncoder
	SamplingConfig   = zap.SamplingConfig
)

func WithLevel(level Level) Option {
	return func(config *zap.Config) {
		config.Level.SetLevel(level)
	}
}

func WithDevelopment(development bool) Option {
	return func(config *zap.Config) {
		config.Development = development
	}
}

func WithDisableCaller(disableCaller bool) Option {
	return func(config *zap.Config) {
		config.DisableCaller = disableCaller
	}
}

func WithDisableStacktrace(disableStacktrace bool) Option {
	return func(config *zap.Config) {
		config.DisableStacktrace = disableStacktrace
	}
}

func WithSampling(sampling *SamplingConfig) Option {
	return func(config *zap.Config) {
		config.Sampling = sampling
	}
}

func WithEncoding(encoding Encoding) Option {
	return func(config *zap.Config) {
		config.Encoding = encoding
	}
}

func WithEncoderMessageKey(encoderMessageKey string) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.MessageKey = encoderMessageKey
	}
}

func WithEncoderLevelKey(encoderLevelKey string) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.LevelKey = encoderLevelKey
	}
}

func WithEncoderTimeKey(encoderTimeKey string) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.TimeKey = encoderTimeKey
	}
}

func WithEncoderNameKey(encoderNameKey string) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.NameKey = encoderNameKey
	}
}

func WithEncoderCallerKey(encoderCallerKey string) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.CallerKey = encoderCallerKey
	}
}

func WithEncoderFunctionKey(encoderFunctionKey string) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.FunctionKey = encoderFunctionKey
	}
}

func WithEncoderStacktraceKey(encoderStacktraceKey string) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.StacktraceKey = encoderStacktraceKey
	}
}

func WithEncoderLineEnding(encoderLineEnding string) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.LineEnding = encoderLineEnding
	}
}

func WithEncoderLevel(encoderLevel LevelEncoder) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.EncodeLevel = encoderLevel
	}
}

func WithEncoderTime(encoderTime TimeEncoder) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.EncodeTime = encoderTime
	}
}

func WithEncoderDuration(encoderDuration DurationEncoder) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.EncodeDuration = encoderDuration
	}
}

func WithEncoderCaller(encoderCaller CallerEncoder) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.EncodeCaller = encoderCaller
	}
}

func WithEncoderName(encoderName NameEncoder) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.EncodeName = encoderName
	}
}

func WithEncoderNewReflectedEncoder(encoderNewReflectedEncoder func(io.Writer) ReflectedEncoder) Option {
	return func(config *zap.Config) {
		config.EncoderConfig.NewReflectedEncoder = encoderNewReflectedEncoder
	}
}

func WithOutputPaths(outputPaths ...string) Option {
	return func(config *zap.Config) {
		config.OutputPaths = outputPaths
	}
}

func WithErrorOutputPaths(errorOutputPaths ...string) Option {
	return func(config *zap.Config) {
		config.ErrorOutputPaths = errorOutputPaths
	}
}

func WithInitialFields(initialFields map[string]interface{}) Option {
	return func(config *zap.Config) {
		config.InitialFields = initialFields
	}
}
