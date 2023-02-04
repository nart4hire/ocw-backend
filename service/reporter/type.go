package reporter

import "context"

type ReporterPayload struct {
	Timestamp string `json:"dt"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

type Reporter interface {
	Send(payload ReporterPayload)
	Flush()
	Start(ctx context.Context)
	Clear()
}
