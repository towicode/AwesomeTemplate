package ports

import cloudevents "github.com/cloudevents/sdk-go/v2"


//	Example interface for a port. Normally it would have more functions than just Create, you may have a full
//	CRUD interface here, or any additional functions for your port.
type AlphaPort interface {
	Create(msg cloudevents.Event) cloudevents.Event
}

//	Example interface for an outgoing port. This one only has one function to store the data.
type StoragePort interface{
	Store(msg cloudevents.Event) error
}