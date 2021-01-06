package domain

import (
	"awesomeTemplate/adapters"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

/*
At the moment this file doesn't do much except connect the receivers and the senders
but exists in case logic needs to happen here for whatever circumstance
*/


type PortProcessor struct {
	Domain Domain
	Storage adapters.StorageAdpater
}

func (c PortProcessor) Create(msg cloudevents.Event) cloudevents.Event{
	return c.Domain.create(msg)
}

func (c PortProcessor) Store(msg cloudevents.Event)  {
	c.Storage.DoAction(msg)
}