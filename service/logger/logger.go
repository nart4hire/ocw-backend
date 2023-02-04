package logger

import (
	"github.com/sirupsen/logrus"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger/hooks"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gitlab.informatika.org/ocw/ocw-backend/utils/log"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

// Object Builder
var createdLogger *LogrusLogger = nil

func New(
	env *env.Environment,
	logUtil log.LogUtils,
	hooks hooks.LogrusHookCollection,
) *LogrusLogger {
	if createdLogger != nil {
		return createdLogger
	}

	log := logrus.New()

	log.SetFormatter(&LogrusFormatter{
		Util: logUtil,
	})

	for _, hook := range hooks {
		if hook.IsProductionOnly && env.AppEnvironment == "PRODUCTION" ||
			!hook.IsProductionOnly {
			log.AddHook(hook.Hook)
		}
	}

	if env.AppEnvironment == "PRODUCTION" {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logrus.DebugLevel)
	}

	createdLogger = &LogrusLogger{
		logger: log,
	}

	return createdLogger
}
