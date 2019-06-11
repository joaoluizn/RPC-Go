package naming

import (
	"fmt"
	"net/http"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/network"
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
func (n *NamingService) RegisterServices(httpRequest *http.Request) []byte {
	response := make([]string, 0)
	registrationReq := n.marshaller.UnmarshalNamingServiceRegistration(httpRequest)
	// log.Printf("Register Request Received from: %s\n", registrationReq.ServerAddress)
	service_list := network.MakeServiceList(registrationReq.ServicesNames, registrationReq.ServerAddress)

	for index := range service_list {
		response = append(response, n.registerService(service_list[index]))
	}

	return n.marshaller.MarshallRegistrationResponse(response)
}

// LookupService gets the first response  for the naming service given
func (n *NamingService) LookupService(serviceName string) []byte {
	var response string

	// log.Printf("Looking up for service: '%s'\n", serviceName)
	_, nameExists := n.registeredRemoteServices[serviceName]

	if !nameExists {
		return n.marshaller.MarshallLookupResponse(response)
	} else {
		response = n.registeredRemoteServices[serviceName].Address

	}

	return n.marshaller.MarshallLookupResponse(response)
}

// registerService registers a new service that is available for clients
func (n *NamingService) registerService(service *network.Service) string {
	var response string

	serviceName := service.Name
	serviceAddr := service.Address
	_, nameExists := n.registeredRemoteServices[service.Name]

	if !nameExists {
		n.registeredRemoteServices[serviceName] = service
		n.showRegisteredServices()

		response = fmt.Sprintf("Service: '%s' IP: '%s' Status: Register Complete\n", serviceName, serviceAddr)
	} else {
		if n.registeredRemoteServices[serviceName].Address == service.Address {
			response = fmt.Sprintf("Service: %s of IP: %s is already registered\n", serviceName, serviceAddr)
		} else {
			response = fmt.Sprintf("Service: %s is already registered by another Address\n", serviceName)
		}
	}

	return response
}

func (n *NamingService) showRegisteredServices() {
	servicesNames := make([]string, 0)
	mapAddrs := make(map[string]string)

	for key := range n.registeredRemoteServices {
		servicesNames = append(servicesNames, key)
		service := n.registeredRemoteServices[key]
		mapAddrs[key] = service.Address
	}
	// log.Printf("(NamingServerStatus)> Registered Services: #%d service(s): %s. Addresses: %s\n\n", len(servicesNames), servicesNames, mapAddrs)
}
