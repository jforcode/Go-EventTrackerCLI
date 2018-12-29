package main

import (
	"encoding/json"
	"time"

	"github.com/jforcode/Go-DeepError"
)

// Event model to represent an event
type Event struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	Note          string      `json:"note"`
	UserCreatedAt time.Time   `json:"created_at"`
	Type          *EventType  `json:"type"`
	Tags          []*EventTag `json:"tags"`
}

// ToJSON returns a printable json representation
func (event *Event) ToJSON() string {
	fn := "ToJSON"

	eventJSON, err := json.MarshalIndent(event, "", "    ")
	if err != nil {
		return deepError.New(fn, "marshal", err).Error()
	}

	return string(eventJSON)
}

// EventType model to represent a type
type EventType struct {
	Value string `json:"value"`
}

// EventTag model to represent a tag
type EventTag struct {
	Value string `json:"value"`
}
