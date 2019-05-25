package storage

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/joaoluizn/RPC-go/network"
)

// NewRemoteService builds a new instance of RemoteService
func NewRemoteService() *RemoteService {
	return &RemoteService{
		// All operations that Storage Service contains
		services:   make(map[string]interface{}),
		marshaller: network.NewMarshaller(),
		clientHttp: &http.Client{
			// 10 seconds until timeout in any request.
			Timeout: 10 * time.Second},
	}
}

// RemoteService holds the services available for clients
type RemoteService struct {
	services   map[string]interface{}
	marshaller *network.Marshaller
	clientHttp *http.Client
}

// RegisterService registers a new service that is available for client
func (r *RemoteService) RegisterService(serviceName string, serviceInstance interface{}) {
	r.services[serviceName] = serviceInstance
}

// GetService gets a service instance
func (r *RemoteService) GetService(serviceName string) interface{} {
	return r.services[serviceName]
}

// getServicesNames gets all services instances names
func (r *RemoteService) getServicesNames() []string {
	names := make([]string, 0)
	for name := range r.services {
		names = append(names, name)
	}
	return names
}

// SaveServicesToNamingService binds services on naming service server
func (r *RemoteService) SaveServicesToNamingService(serviceAddr string, namingServerAddr string) {
	// dialer := network.GetTCPDialer(namingServerAddr)
	// defer dialer.Close()

	// Get all Storage Operations
	names := r.getServicesNames()
	// log.Printf(internal.MsgRegisteringService, names, address)

	namingServiceRegistration := network.NewNamingServiceRegistration(names, serviceAddr)
	namingServiceRegistrationBytes := r.marshaller.MarshalNamingServiceRegistration(namingServiceRegistration)

	r.Register(namingServerAddr, namingServiceRegistrationBytes)
	// WatchNamingService(dialer)
}

// Send sends a invoke request to remote service
func (r *RemoteService) Register(namingServerAddr string, request *bytes.Buffer) *http.Response {
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

	return response
}
