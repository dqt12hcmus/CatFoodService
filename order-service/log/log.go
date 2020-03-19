package log

import (
	log "github.com/sirupsen/logrus"
)

func SetupLog(service string) func() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)
	return func() {}
}
