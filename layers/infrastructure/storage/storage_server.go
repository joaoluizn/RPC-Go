package storage

import (
	"log"
	"net"
	"net/http"

	"github.com/joaoluizn/RPC-go/network"
)

func NewStorageServiceServer(serviceAddr string, servicePort string) *StorageServiceServer {
	return &StorageServiceServer{
		registrationServerAddress: serviceAddr,
		port:                      servicePort,
		requestHandler:            NewStorageServiceRequestHandler(),
	}
}

// RemoteServiceServer holds services that can be invoke for clients
type StorageServiceServer struct {
	registrationServerAddress string
	port                      string
	requestHandler            *StorageServiceRequestHandler
}

// Run runs the remote service
func (r *StorageServiceServer) Run() {
	listener, address := network.GetTCPListener(r.port)
	go r.runHTTPServerForServicesInvocation(listener, address)
	// r.bindServicesInNamingService(address)
}

// RegisterServiceInNamingService adds a new service that will be available for clients
func (r *StorageServiceServer) RegisterServiceInNamingService(name string, instance interface{}) {
	r.requestHandler.Invoker.RemoteService.RegisterService(name, instance)
}

// runHTTPServerForServicesInvocation brings up the http server that handles services invoke requests
func (r *StorageServiceServer) runHTTPServerForServicesInvocation(listener net.Listener, address string) {
	// log.Printf(internal.MsgRunningServicesInvoke, address)
	http.HandleFunc("/invoke/", r.requestHandler.HandleInvokeRequest)

	errServe := http.Serve(listener, nil)
	if errServe != nil {
		log.Fatal(errServe.Error())
	}
}

// // bindServicesInNamingService binds services on naming service server
// func (r *StorageServiceServer) bindServicesInNamingService(address string) {
// 	r.requestHandler.Invoker.RemoteService.BindNamingService(address, r.registrationServerAddress)
// }
