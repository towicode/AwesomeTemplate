package adapters

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type StorageAdpater struct {
}


func (c StorageAdpater) Store(event cloudevents.Event)  {
	println(event.Data())
}