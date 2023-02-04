package log

type LogUtils interface {
	PrintFormattedOutput(text string, process string, logType string, color Color)
	FormattedOutput(text string, process string, logType string, color Color) string
	ColoredOutput(text string, color Color) string
}
