package log

type JsonMessage struct {
	// Timestamp string `json:"timestamp"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

type Transport interface {
	log(json *JsonMessage) error
}

type Logger struct {
	transports []Transport
}

type LoggerOptions struct {
	Transports []Transport
}
