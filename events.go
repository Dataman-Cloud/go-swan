package swan

import (
	"github.com/Dataman-Cloud/swan/types"
)

// EventsChannel is a channel to receive events upon
type EventsChannel chan *Event

type Event struct {
	ID    string
	Event string
	Data  interface{}
}

func GetEvent(eventType string) (*Event, error) {
	event := new(Event)
	event.Data = new(types.TaskEvent)
	return event, nil
}
