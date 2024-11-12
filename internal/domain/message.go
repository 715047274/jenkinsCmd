package domain

import "time"

type Message struct {
	ID        int64
	EventType string
	Payload   string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	StatusPending    = "pending"
	StatusProcessing = "processing"
	StatusCompleted  = "completed"
	StatusFailed     = "failed"
)
