package main

import "time"

// Event model to represent an event
type Event struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	Note          string      `json:"note"`
	UserCreatedAt time.Time   `json:"created_at"`
	Type          *EventType  `json:"type"`
	Tags          []*EventTag `json:"tags"`
}

// EventType model to represent a type
type EventType struct {
	Value string `json:"value"`
}

// EventTag model to represent a tag
type EventTag struct {
	Value string `json:"value"`
}
