package dom

type Event struct{} //fix this up later

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
