package services

import (
	"awesomeTemplate/ports"
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type service struct {
	storage ports.StoragePort
}


func (srv *service) Create (event cloudevents.Event) cloudevents.Event{

	srv.storage.Store(event)

	ok :=  cloudevents.NewEvent()
	ok.SetSource("example/uri")
	ok.SetType("example.type")
	natsJson, _:= json.Marshal(string("Okay 200"))
	ok.SetData(cloudevents.ApplicationJSON, natsJson)
	return ok
}