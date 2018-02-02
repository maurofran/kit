package domain

import "time"

// Event is the interface exposed by domain events.
type Event interface {
	Type() string
	OccurredOn() time.Time
	Version() int
}
