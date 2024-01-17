package log_test

import (
	stdlog "log"
	"testing"
	"time"

	"github.com/hachimi-lab/log"
	"github.com/hachimi-lab/log/aliyun"
	"github.com/hachimi-lab/rotatelogs"
)

func TestLog(t *testing.T) {
	logger, err := log.New(
		log.WithDevelopment(false),
		log.WithLevel(log.InfoLevel),
		log.WithDisableStacktrace(false),
	).Extend(
		rotatelogs.New(
			"./logs/hachimi-lab.log",
			rotatelogs.WithTimePeriod(rotatelogs.Daily),
			rotatelogs.WithMaxAge(time.Hour*24*7),
		),
	).Extend(
		aliyunlog.New(aliyunlog.Config{}),
		log.ErrorLevel,
	).Build()
	if err != nil {
		stdlog.Fatal(err)
	}
	log.SetLogger(logger.Named("hachimi-lab"))

	log.Info("hello world", log.String("name", "Kaka"))
}
