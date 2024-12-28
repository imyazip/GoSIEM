package models

type SecurityEvent struct {
	ID               int32
	LogID            int32
	EventType        string
	EventDescription string
	DetectedAt       string
	CreatedAt        string
}

type Rule struct {
	ID   int64  `json:"id"`
	JSON string `json:"rule"`
}
