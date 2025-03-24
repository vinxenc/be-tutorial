package log

func NewLogger(options *LoggerOptions) *Logger {
	return &Logger{
		transports: options.Transports,
	}
}

func (l *Logger) log(level, message string) {
	json := &JsonMessage{
		Level:   level,
		Message: message,
	}

	for _, transport := range l.transports {
		_ = transport.log(json)
	}
}

func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

func (l *Logger) Warn(message string) {
	l.log("WARN", message)
}

func (l *Logger) Debug(message string) {
	l.log("DEBUG", message)
}
