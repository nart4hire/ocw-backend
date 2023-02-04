package log

import (
	"fmt"
	"time"
)

type LogUtilsImpl struct{}

func (l LogUtilsImpl) FormattedOutput(text string, process string, logType string, color Color) string {
	return fmt.Sprintf("%s %s: [%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05 MST"),
		l.ColoredOutput(logType, color),
		process,
		text,
	)
}

func (LogUtilsImpl) ColoredOutput(text string, color Color) string {
	return fmt.Sprintf("%s%s%s",
		color,
		text,
		Reset,
	)
}

func (l LogUtilsImpl) PrintFormattedOutput(text string, process string, logType string, color Color) {
	print(
		l.FormattedOutput(text, process, logType, color),
	)
}
