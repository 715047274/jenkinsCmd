package domain

type Event interface {
	EventType() string
}
