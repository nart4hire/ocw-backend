package hooks

import (
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.informatika.org/ocw/ocw-backend/service/reporter"
)

type LogrusReporter struct {
	Reporter reporter.Reporter
}

func (LogrusReporter) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
	}
}

func (l LogrusReporter) Fire(entry *logrus.Entry) error {
	payload := reporter.ReporterPayload{
		Level:     entry.Level.String(),
		Timestamp: entry.Time.Format(time.RFC3339),
		Message:   entry.Message,
	}

	l.Reporter.Send(payload)
	return nil
}
