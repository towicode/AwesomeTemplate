package adapters

import (
	"awesomeTemplate/core/services"
	"awesomeTemplate/ports"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/go-nats"
)

type AlphaAdapter struct {
	Port ports.AlphaPort
}


//	Do action is a function that takes a Nats Message and then passes it to the port to be processed, in this
//	example I have it returning a cloud event but it may return nothing at all.
func (c AlphaAdapter) doAction(msg *nats.Msg) cloudevents.Event {
	// The do action should determine the
	// proper course of action for the NATS message
	// and process it.domain


	//Here you would have some sort of logic to convert a
	//Nats message to a cloudevent with the expected format
	//for the port. However... for this example we just
	//pretend that has happened
	viz := services.Nats2CloudEvent{}
	ce:= viz.Convert(msg)
	returnEvent := c.Port.Create(ce)

	return returnEvent

}


//	This function tells the adapter to start listening for updates. In this case the NATS subscriber would
//	listen to a message and then call doAction whenever a message is received.
func (c AlphaAdapter) StartSubscriber() {
	// Logic to start NATS subscriber would live here
	// See cred microservice as example
	incomingMessage := nats.Msg{}
	c.doAction(&incomingMessage)

}
