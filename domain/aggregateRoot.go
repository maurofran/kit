package domain

// AggregateRoot is the parent struct used to manage domain events.
type AggregateRoot struct {
	events []Event
}

// AndEventsFrom will append all the events from supplied aggregate root to the receiver, returning it.
func (ar *AggregateRoot) AndEventsFrom(other *AggregateRoot) *AggregateRoot {
	events := append(ar.events, other.events...)
	ar.events = events
	return ar
}

// AndEvent will register the supplied event to the receiver aggregate root, returning it
func (ar *AggregateRoot) AndEvent(event Event) *AggregateRoot {
	ar.RegisterEvent(event)
	return ar
}

// RegisterEvent will queue a new event to this aggregate root, returning the created event.
func (ar *AggregateRoot) RegisterEvent(event Event) Event {
	events := append(ar.events, event)
	ar.events = events
	return event
}

// DomainEvents will return a copy of internal events of receiver aggregate root.
func (ar *AggregateRoot) DomainEvents() []Event {
	var events []Event
	copy(events, ar.events)
	return events
}

// ClearDomainEvents will clear the domain events of this aggregate root.
func (ar *AggregateRoot) ClearDomainEvents() {
	ar.events = nil
}
