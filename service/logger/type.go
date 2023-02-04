package logger

type Logger interface {
	Debug(message string)
	Info(message string)
	Warning(message string)
	Error(message string)
}
