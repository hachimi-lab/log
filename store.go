package log

import (
	"time"

	"github.com/hachimi-lab/rotatelogs"
)

type (
	StoreOption = rotatelogs.Option
	RotateTime  = rotatelogs.RotateTime
)

const (
	StoreEveryMinute = rotatelogs.EveryMinute
	StoreEveryHour   = rotatelogs.EveryHour
	StoreEveryDay    = rotatelogs.EveryDay
)

func StoreWithRotateTime(rotateTime RotateTime) StoreOption {
	return rotatelogs.WithRotateTime(rotateTime)
}

func StoreWithMaxAge(maxAge time.Duration) StoreOption {
	return rotatelogs.WithMaxAge(maxAge)
}
