package domain

import (
	"awesomeTemplate/ports"
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)


type Domain struct {
	storagePort ports.StoragePort
}

func (m Domain) SetStoragePort(port PortProcessor) {
	m.storagePort = port
}


// This is the actual core logic of the process, here the program takes cloud events and does whatever with them,
// processing them and passing them onto whatever port/adapter they need to goto.
func (m Domain) create(msg cloudevents.Event) cloudevents.Event {

	event :=  cloudevents.NewEvent()
	event.SetSource("example/uri")
	event.SetType("example.type")
	natsJson, _:= json.Marshal(string("abc123"))
	event.SetData(cloudevents.ApplicationJSON, natsJson)

	return event
}

func (m Domain) storeVitalInformation(msg cloudevents.Event) error {
	return m.storagePort.Store(msg)
}