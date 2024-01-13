package log

import (
	"time"

	"github.com/hachimi-lab/rotatelogs"
)

type (
	RotateType  = rotatelogs.RotateType
	StoreOption = rotatelogs.Option
	RotateTime  = rotatelogs.RotateType
)

const (
	RotateEveryMinute = rotatelogs.EveryMinute
	RotateEveryHour   = rotatelogs.EveryHour
	RotateEveryDay    = rotatelogs.EveryDay
)

func StoreWithRotateType(rotateType RotateType) StoreOption {
	return rotatelogs.WithRotateType(rotateType)
}

func StoreWithMaxAge(maxAge time.Duration) StoreOption {
	return rotatelogs.WithMaxAge(maxAge)
}
