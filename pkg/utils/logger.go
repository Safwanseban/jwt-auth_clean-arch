package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetLogger() zerolog.Logger {

	thislogger := log.With().Logger()
	return thislogger
}
