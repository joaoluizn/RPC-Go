package storage

import (
	"log"
	"net"
	"net/http"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/network"
)

// NewStorageServiceServer create StorageServiceServer Entity
func NewStorageServiceServer(storageServerAddr string, namingServerAddr string) *StorageServiceServer {
	return &StorageServiceServer{
		namingServerAddr:  namingServerAddr,
		storageServerAddr: storageServerAddr,
		requestHandler:    NewStorageServiceRequestHandler(),
	}
}

// StorageServiceServer Entity Responsible to save services that can be invoked from clients;
type StorageServiceServer struct {
	namingServerAddr  string
	storageServerAddr string
	requestHandler    *StorageServiceRequestHandler
}

// RunRemoteServer core function to run the Remote Server;
func (r *StorageServiceServer) RunRemoteServer() {
	listener := network.GetTCPListener(r.storageServerAddr)
	r.bindServicesToNamingService(r.storageServerAddr)
	r.runHTTPServerForServicesInvocation(listener, r.storageServerAddr)
}

// RegisterServiceInLocalStorage adds a new service that will be available for clients;
func (r *StorageServiceServer) RegisterServiceInLocalStorage(name string, instance interface{}) {
	r.requestHandler.Invoker.RemoteService.RegisterService(name, instance)
}

// runHTTPServerForServicesInvocation Run a http server for the /invoke endpoint
func (r *StorageServiceServer) runHTTPServerForServicesInvocation(listener net.Listener, storageServerAddr string) {
	log.Printf("Running Storage Remote Server on: %s\n\n", storageServerAddr)

	http.HandleFunc("/invoke/", r.requestHandler.HandleInvokeRequest)

	errServe := http.Serve(listener, nil)
	if errServe != nil {
		log.Fatal(errServe.Error())
	}
}

// bindServicesInNamingService binds new services to Naming Service
func (r *StorageServiceServer) bindServicesToNamingService(serviceAddr string) {
	log.Printf("Binding Services from: %s to Naming Service.\n", serviceAddr)
	r.requestHandler.Invoker.RemoteService.SaveServicesToNamingService(serviceAddr, r.namingServerAddr)
}
