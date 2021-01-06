package services

import (
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/go-nats"
)

type Nats2CloudEvent struct {
}

func (Nats2CloudEvent) Convert(msg *nats.Msg) cloudevents.Event {
	event :=  cloudevents.NewEvent()
	event.SetSource("example/uri")
	event.SetType("example.type")
	natsJson, _:= json.Marshal(string(msg.Data))
	event.SetData(cloudevents.ApplicationJSON, natsJson)
	return event
}
