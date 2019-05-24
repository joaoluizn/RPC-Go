package storage

import (
	"github.com/joaoluizn/RPC-go/network"
)

// NewRemoteService builds a new instance of RemoteService
func NewRemoteService() *RemoteService {
	return &RemoteService{
		// All operations that Storage Service contains
		services:   make(map[string]interface{}),
		marshaller: network.NewMarshaller(),
	}
}

// RemoteService holds the services available for clients
type RemoteService struct {
	services   map[string]interface{}
	marshaller *network.Marshaller
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

// // BindNamingService binds services on naming service server
// func (r *RemoteService) BindNamingService(address string, namingServiceAddress string) {
// 	dialer := network.GetTCPDialer(namingServiceAddress)
// 	defer dialer.Close()

// 	names := r.getServicesNames()
// 	// log.Printf(internal.MsgRegisteringService, names, address)

// 	namingServiceRegistration := request.NewNamingServiceRegistration(names, address)
// 	namingServiceRegistrationBytes := r.marshaller.MarshalNamingServiceRegistration(namingServiceRegistration)
// 	err := json.NewEncoder(dialer).Encode(namingServiceRegistrationBytes)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	WatchNamingService(dialer)
// }
