package log_test

import (
	stdlog "log"
	"testing"
	"time"

	"github.com/hachimi-lab/log"
)

func TestLog(t *testing.T) {
	logger, err := log.New(
		log.WithDevelopment(false),
		log.WithLevel(log.InfoLevel),
		log.WithDisableStacktrace(false),
	).Store(
		"./logs/hachimi-lab.log",
		log.StoreWithTimePeriod(log.Daily),
		log.StoreWithMaxAge(time.Hour*24*7),
	).Build()
	if err != nil {
		stdlog.Fatal(err)
	}
	log.SetLogger(logger.Named("hachimi-lab"))

	log.Info("hello world", log.String("name", "Kaka"))
}
