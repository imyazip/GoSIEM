package models

import "time"

type SecurityEvent struct {
	ID               int64     `json:"id"`
	EventType        string    `json:"event_type"`
	EventDescription string    `json:"event_description"`
	DetectedAt       time.Time `json:"detected_at"`
	ReadFlag         bool      `json:"read_flag"`
}

type Rule struct {
	ID   int64  `json:"id"`
	JSON string `json:"rule"`
}
