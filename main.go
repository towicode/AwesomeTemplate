package awesomeTemplate

import (
	"awesomeTemplate/adapters"
	"awesomeTemplate/core/domain"
)

/*
This is a template for a very basic microservice that only contains
one function 'create'.

the microservice has two adapters, one incoming and one outgoing.

StorageAdpater represents the incoming message, in this case the
message is coming from a NATS queue.
*/

func main() {

	// Implementation follows a right to left process for the hexagon
	sa := adapters.StorageAdapter{}
	dmain := domain.Domain{
		StoragePort: sa,
	}
	aa := adapters.AlphaAdapter{
		Port: dmain,
	}
	go aa.StartSubscriber()
}
