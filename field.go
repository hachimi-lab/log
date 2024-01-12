package log

import (
	"time"

	"go.uber.org/zap"
)

type Field = zap.Field

var (
	Any         = zap.Any
	Bool        = zap.Bool
	BoolP       = zap.Boolp
	Int         = zap.Int
	IntP        = zap.Intp
	Int8        = zap.Int8
	Int8P       = zap.Int8p
	Int16       = zap.Int16
	Int16P      = zap.Int16p
	Int32       = zap.Int32
	Int32P      = zap.Int32p
	Int64       = zap.Int64
	Int64P      = zap.Int64p
	Uint        = zap.Uint
	UintP       = zap.Uintp
	Uintptr     = zap.Uintptr
	UintptrP    = zap.Uintptrp
	Uint8       = zap.Uint8
	Uint8P      = zap.Uint8p
	Uint16      = zap.Uint16
	Uint16P     = zap.Uint16p
	Uint32      = zap.Uint32
	Uint32P     = zap.Uint32p
	Uint64      = zap.Uint64
	Uint64P     = zap.Uint64p
	Float32     = zap.Float32
	Float32P    = zap.Float32p
	Float64     = zap.Float64
	Float64P    = zap.Float64p
	Complex64   = zap.Complex64
	Complex64P  = zap.Complex64p
	Complex128  = zap.Complex128
	Complex128P = zap.Complex128p
	String      = zap.String
	StringP     = zap.Stringp
	ByteString  = zap.ByteString
	Binary      = zap.Binary
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	TimeP       = zap.Timep
	Skip        = zap.Skip
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = func(key string, val time.Duration) Field { return String(key, val.String()) }
	DurationP   = func(key string, val *time.Duration) Field { str := (*val).String(); return StringP(key, &str) }
	Object      = zap.Object
	Inline      = zap.Inline
	Err         = zap.Error
)
