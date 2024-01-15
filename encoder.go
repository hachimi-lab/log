package log

import (
	"time"

	"github.com/hachimi-lab/rotatelogs"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Encoder struct {
	internal zapcore.Encoder
	cores    []zapcore.Core
	config   *zap.Config
	err      error
}

func New(cfg ...Config) *Encoder {
	config := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			TimeKey:       "ts",
			NameKey:       "name",
			CallerKey:     "caller",
			StacktraceKey: "stack",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
			},
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	for _, v := range cfg {
		v(config)
	}

	if len(config.Encoding) == 0 {
		if config.Development {
			config.Encoding = ConsoleEncoding
		} else {
			config.Encoding = JsonEncoding
		}
	}

	encoder := &Encoder{config: config}
	switch config.Encoding {
	case ConsoleEncoding:
		encoder.internal = zapcore.NewConsoleEncoder(config.EncoderConfig)
	case JsonEncoding:
		encoder.internal = zapcore.NewJSONEncoder(config.EncoderConfig)
	default:
		encoder.err = errors.Errorf("unknown encoding: %s", config.Encoding)
	}

	return encoder
}

func (slf *Encoder) Store(logPath string, opts ...StoreOption) *Encoder {
	if slf.err != nil {
		return slf
	}
	writer, err := rotatelogs.New(logPath, opts...)
	if err != nil {
		slf.err = err
		return slf
	}
	slf.cores = append(slf.cores,
		zapcore.NewCore(slf.internal, NewWriteSyncer(writer), slf.config.Level.Level()),
	)
	return slf
}

func (slf *Encoder) Extend(writeSyncer WriteSyncer, levelEnabler ...LevelEnabler) *Encoder {
	if slf.err != nil {
		return slf
	}
	if writeSyncer == nil {
		return slf
	}
	var enabler LevelEnabler
	enabler = slf.config.Level
	if len(levelEnabler) > 0 {
		enabler = levelEnabler[0]
	}
	slf.cores = append(slf.cores, zapcore.NewCore(slf.internal, writeSyncer, enabler))
	return slf
}

func (slf *Encoder) Build(options ...Option) (*Logger, error) {
	if slf.err != nil {
		return nil, slf.err
	}
	ins, err := slf.config.Build()
	if err != nil {
		return nil, err
	}
	options = append([]zap.Option{BuildWithAddCaller(), BuildWithAddCallerSkip(1)}, options...)
	options = append(options, BuildWithWrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(append(slf.cores, core)...)
	}))
	ins = ins.WithOptions(options...)
	return &Logger{ins}, nil
}
