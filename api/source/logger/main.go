package logger

import (
	"github.com/rs/zerolog"
	dl "github.com/rs/zerolog/log"
	"os"
)

var (
	Log = zerolog.Logger{}
)

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true}
	Log = zerolog.New(output).Hook(SeverityHook{}).With().Timestamp().Logger()
	dl.Level(zerolog.DebugLevel)
}

type SeverityHook struct{}

func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		e.Str("severity", level.String())
	}
}
