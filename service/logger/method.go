package logger

func (l *LogrusLogger) Debug(message string) {
	l.logger.Debug(message)
}

func (l *LogrusLogger) Info(message string) {
	l.logger.Info(message)
}

func (l *LogrusLogger) Warning(message string) {
	l.logger.Warn(message)
}

func (l *LogrusLogger) Error(message string) {
	l.logger.Error(message)
}
