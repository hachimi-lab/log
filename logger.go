package log

var (
	logger *Logger
)

func init() {
	logger, _ = New().Build()
}

func SetLogger(l *Logger) {
	if l == nil {
		return
	}
	_ = l.internal.Sync()
	logger = l
}

func GetLogger() *Logger {
	return logger
}

func Debug(msg string, fields ...Field) {
	logger.internal.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	logger.internal.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	logger.internal.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	logger.internal.Error(msg, fields...)
}

func DPanic(msg string, fields ...Field) {
	logger.internal.DPanic(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	logger.internal.Panic(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	logger.internal.Fatal(msg, fields...)
}
