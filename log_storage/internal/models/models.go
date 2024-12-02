package models

type SecurityEvent struct {
	ID               int32
	LogID            int32
	EventType        string
	EventDescription string
	DetectedAt       string
	CreatedAt        string
}
