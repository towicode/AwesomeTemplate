package adapters

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type StorageAdapter struct {
}


//	This is an example of a Driven adapter. The core will call this logic. This is normally
//	Where you might see a SQL database or something similar
func (c StorageAdapter) Store(event cloudevents.Event) error{
	println("We reached the storage adapter!")
	return nil
}