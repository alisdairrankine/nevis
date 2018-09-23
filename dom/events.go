package dom

// Events  aren't really used atm

// Event represents a DOM Event
type Event struct{}

// EventHandler is an event triggered  function
type EventHandler func(e Event)

type RegisteredEventHandler struct {
	eventIdentifier string
	eventHandler    EventHandler
}

func CreateEventRegistration(eventName string) func(EventHandler) RegisteredEventHandler {
	return func(handler EventHandler) RegisteredEventHandler {
		return RegisteredEventHandler{
			eventIdentifier: eventName,
			eventHandler:    handler,
		}
	}
}
