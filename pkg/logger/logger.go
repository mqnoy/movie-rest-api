package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	once   sync.Once
	logger zerolog.Logger
)

func Logger() *zerolog.Logger {
	return &logger
}

func Info() *zerolog.Event {
	return Logger().Info()
}

func Error() *zerolog.Event {
	return Logger().Error()
}

func Fatal() *zerolog.Event {
	return Logger().Fatal()
}

func Debug() *zerolog.Event {
	return Logger().Debug()
}

func Panic() *zerolog.Event {
	return Logger().Panic()
}

func Trace() *zerolog.Event {
	return Logger().Trace()
}

func Warn() *zerolog.Event {
	return Logger().Warn()
}

func init() {
	once.Do(func() {
		writer := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("drop logs %d", missed)
		})

		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

		logger = zerolog.New(writer).
			With().
			Timestamp().
			Logger()

	})
}
