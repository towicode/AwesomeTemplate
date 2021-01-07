package domain

import cloudevents "github.com/cloudevents/sdk-go/v2"

type Domain struct {
}

func (c Domain) VerifyMessage(event cloudevents.Event) bool {
	return true
}
