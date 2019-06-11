package storage

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/network"
)

// NewRemoteService builds RemoteService Entity
func NewRemoteService() *RemoteService {
	return &RemoteService{
		// All services that this remote server contains.
		services:   make(map[string]interface{}),
		marshaller: network.NewMarshaller(),
		clientHttp: &http.Client{
			// 10 seconds until timeout in any request.
			Timeout: 10 * time.Second},
	}
}

// RemoteService Entity responsible to store services available to clients
type RemoteService struct {
	services   map[string]interface{}
	marshaller *network.Marshaller
	clientHttp *http.Client
}

// RegisterService Register a locally service to future bind to Naming Service
func (r *RemoteService) RegisterService(serviceName string, serviceInstance interface{}) {
	r.services[serviceName] = serviceInstance
}

// GetService Get a service entity
func (r *RemoteService) GetService(serviceName string) interface{} {
	return r.services[serviceName]
}

// getServicesNames Get all services names
func (r *RemoteService) getServicesNames() []string {
	names := make([]string, 0)
	for name := range r.services {
		names = append(names, name)
	}
	return names
}

// SaveServicesToNamingService binds services on naming service server
func (r *RemoteService) SaveServicesToNamingService(serviceAddr string, namingServerAddr string) {
	// Get all Services for this Remote Server
	servicesNames := r.getServicesNames()

	namingServiceRegistration := network.NewNamingServiceRegistration(servicesNames, serviceAddr)
	namingServiceRegistrationBytes := r.marshaller.MarshalNamingServiceRegistration(namingServiceRegistration)

	log.Printf("Registering Services: %s ~from: %s\n\n", servicesNames, serviceAddr)
	response := r.Register(namingServerAddr, namingServiceRegistrationBytes)
	for index := range response {
		log.Printf("ServiceRegistrationStatus: %s", response[index])
	}
}

// Register sends a registration request to Naming Server
func (r *RemoteService) Register(namingServerAddr string, request *bytes.Buffer) []string {
	response, err := r.clientHttp.Post(
		// URL
		"http://"+namingServerAddr+"/register/",
		// ContentType
		"service/json",
		// Data
		request,
	)
	if err != nil {
		log.Fatal("storage_service (Register): ", err.Error())
	}

	return r.marshaller.UnMarshallRegistrationResponse(response)
}
