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

	//	First we have to create the domain
	myDomain := domain.Domain{}

	//	Then we create the outgoing adapters
	myOutgoingAdapter := adapters.StorageAdpater{}

	//	Then we create the port processor
	myPortProcessor := domain.PortProcessor{
		Domain:  myDomain,
		Storage: myOutgoingAdapter,
	}
	//	Then we retro-set the port processor to the domain, so the domain can access outoging ports
	myDomain.SetStoragePort(myPortProcessor)

	//	Then we create incoming adapters so incoming messages can access the domain
	myAdapter := adapters.AlphaAdapter{
		Port: myPortProcessor,
	}

	//	Finally we start all incoming Adapters
	go myAdapter.StartSubscriber()

	// some sort of logic to keep the program from
	// ending
}
