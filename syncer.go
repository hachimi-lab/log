package log

import (
	"io"
	"os"

	"go.uber.org/zap/zapcore"
)

type WriteSyncer = zapcore.WriteSyncer

func NewWriteSyncer(ws ...io.Writer) WriteSyncer {
	if len(ws) == 0 {
		return nil
	}
	multi := make([]WriteSyncer, len(ws))
	for i, w := range ws {
		multi[i] = zapcore.AddSync(w)
	}
	return zapcore.NewMultiWriteSyncer(multi...)
}

func StdoutWriteSyncer() WriteSyncer {
	return zapcore.AddSync(io.Writer(os.Stdout))
}
