package log

import (
	"io"
	"time"

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
			FunctionKey:   "func",
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

func (slf *Encoder) Extend(writer io.Writer, lvlEnabler ...LevelEnabler) *Encoder {
	if slf.err != nil || writer == nil {
		return slf
	}
	var enabler LevelEnabler
	enabler = slf.config.Level
	if len(lvlEnabler) > 0 {
		enabler = lvlEnabler[0]
	}
	slf.cores = append(slf.cores,
		zapcore.NewCore(slf.internal, zapcore.AddSync(writer), enabler),
	)
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
