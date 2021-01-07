package domain

import (
	"awesomeTemplate/ports"
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)


type Domain struct {
	StoragePort ports.StoragePort
}


// This is the actual core logic of the process, here the program takes cloud events and does whatever with them,
// processing them and passing them onto whatever port/adapter they need to goto.
func (m Domain) Create(msg cloudevents.Event) cloudevents.Event {

	_ = m.storeVitalInformation(msg)


	//Now that we have stored that important piece of information, we need to send back a message to
	//the sender. In this example we will send an OK message back

	event :=  cloudevents.NewEvent()
	event.SetSource("example/uri")
	event.SetType("example.type")
	natsJson, _:= json.Marshal(string("OK"))
	_ = event.SetData(cloudevents.ApplicationJSON, natsJson)

	return event
}

func (m Domain) storeVitalInformation(msg cloudevents.Event) error {
	return m.StoragePort.Store(msg)
}