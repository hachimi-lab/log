package log

import (
	"time"

	"github.com/hachimi-lab/rotatelogs"
)

type (
	TimePeriod  = rotatelogs.TimePeriod
	StoreOption = rotatelogs.Option
)

const (
	Minutely = rotatelogs.Minutely
	Hourly   = rotatelogs.Hourly
	Daily    = rotatelogs.Daily
)

func StoreWithTimePeriod(timePeriod TimePeriod) StoreOption {
	return rotatelogs.WithTimePeriod(timePeriod)
}

func StoreWithMaxAge(maxAge time.Duration) StoreOption {
	return rotatelogs.WithMaxAge(maxAge)
}
