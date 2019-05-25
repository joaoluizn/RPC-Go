package naming

import (
	"log"
	"net/http"

	"github.com/joaoluizn/RPC-go/network"
)

// NewNamingService builds a new instance of NamingService
func NewNamingService() *NamingService {
	return &NamingService{
		registeredRemoteServices: make(map[string][]*network.Service),
		marshaller:               network.NewMarshaller(),
	}
}

// NamingService is a naming service who holds the urls to all services available for client
type NamingService struct {
	registeredRemoteServices map[string][]*network.Service
	marshaller               *network.Marshaller
}

// RegisterServices registers new services that are available for client
func (n *NamingService) RegisterServices(httpRequest *http.Request) {
	registrationReq := n.marshaller.UnmarshalNamingServiceRegistration(httpRequest)

	log.Printf("Register Request Received from: %s", registrationReq.ServerAddress)

	service_list := network.MakeServiceList(registrationReq.ServicesNames, registrationReq.ServerAddress)

	for index := range service_list {
		log.Printf("%d: %s", index, service_list[index])
		// n.registerService(remoteServiceEntries[index])
	}
}

// TODO: If really needed to keep the watcher, try to implement it with a Dialer using server addr.
