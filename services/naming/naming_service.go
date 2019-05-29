package naming

import (
	"log"
	"net/http"

	"github.com/joaoluizn/RPC-go/network"
)

// NewNamingService builds a new instance of NamingService
func NewNamingService() *NamingService {
	return &NamingService{
		registeredRemoteServices: make(map[string]*network.Service),
		marshaller:               network.NewMarshaller(),
	}
}

// NamingService is a naming service who holds the urls to all services available for client
type NamingService struct {
	registeredRemoteServices map[string]*network.Service
	marshaller               *network.Marshaller
}

// RegisterServices registers new services that are available for client
func (n *NamingService) RegisterServices(httpRequest *http.Request) {
	response := make([]string, 0)
	registrationReq := n.marshaller.UnmarshalNamingServiceRegistration(httpRequest)
	log.Printf("Register Request Received from: %s\n", registrationReq.ServerAddress)
	service_list := network.MakeServiceList(registrationReq.ServicesNames, registrationReq.ServerAddress)

	for index := range service_list {
		response = append(response, n.registerService(service_list[index]))
	}
}

// LookupService gets the first response  for the naming service given
func (n *NamingService) LookupService(serviceName string) []byte {
	var response string

	log.Printf("Looking up for service: '%s'\n", serviceName)
	_, nameExists := n.registeredRemoteServices[serviceName]

	if !nameExists {
		return n.marshaller.MarshallLookupResponse(response)
	} else {
		response = n.registeredRemoteServices[serviceName].Address
	}

	return n.marshaller.MarshallLookupResponse(response)
}

// registerService registers a new service that is available for clients
func (n *NamingService) registerService(service *network.Service) {
	serviceName := service.Name
	serviceAddr := service.Address
	_, nameExists := n.registeredRemoteServices[service.Name]

	if !nameExists {
		log.Printf("Service: '%s' IP: '%s' Status: Register Complete\n", serviceName, serviceAddr)
		n.registeredRemoteServices[serviceName] = service
		n.showRegisteredServices()

		return
	} else {
		if n.registeredRemoteServices[serviceName].Address == service.Address {
			log.Printf("Service: %s of IP: %s is already registered", serviceName, serviceAddr)
		} else {
			log.Printf("Service: %s is already registered by another Address", serviceName)
		}

		// TODO: implement here routine to adding a service that already exist.

	}

	// if !n.addressExists(entry.Name, entry.Address) {
	// 	log.Printf(internal.MsgRegisteringService, entry.Name, entry.Address)
	// 	entries := n.remoteServicesEntries[entry.Name]
	// 	n.remoteServicesEntries[entry.Name] = append(entries, entry)
	// 	n.status()
	// }
}

func (n *NamingService) showRegisteredServices() {
	servicesNames := make([]string, 0)
	mapAddrs := make(map[string]string)

	for key := range n.registeredRemoteServices {
		servicesNames = append(servicesNames, key)
		service := n.registeredRemoteServices[key]
		mapAddrs[key] = service.Address
	}
	log.Printf("(NamingServerStatus)> Registered Services: #%d service(s): %s. Addresses: %s\n\n", len(servicesNames), servicesNames, mapAddrs)
}
