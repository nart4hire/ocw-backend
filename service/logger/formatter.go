package logger

import (
	"github.com/sirupsen/logrus"
	"gitlab.informatika.org/ocw/ocw-backend/utils/log"
)

type LogrusFormatter struct {
	Util log.LogUtils
}

var colorMap = map[logrus.Level]log.Color{
	logrus.TraceLevel: log.ForeWhite,
	logrus.DebugLevel: log.ForeWhite,
	logrus.InfoLevel:  log.ForeGreen,
	logrus.WarnLevel:  log.ForeYellow,
	logrus.ErrorLevel: log.ForeRed,
	logrus.PanicLevel: log.ForeRed,
}

func (l *LogrusFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(l.Util.FormattedOutput(
		entry.Message,
		"App",
		entry.Level.String(),
		colorMap[entry.Level],
	)), nil
}
