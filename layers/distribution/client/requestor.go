package client

import (
	"bytes"
	"net/http"

	"github.com/joaoluizn/RPC-go/layers/infrastructure/client"
	infrastructure "github.com/joaoluizn/RPC-go/layers/infrastructure/pool"
	"github.com/joaoluizn/RPC-go/network"
)

// NewRequestor: Create a Requestor instance;
func NewRequestor(namingServerAddress string) *Requestor {
	return &Requestor{
		namingServerAddress: namingServerAddress,
		requestHandler:      client.NewClientRequestHandler(),
		marshaller:          network.NewMarshaller(),
		namingPool:          infrastructure.NewNamingPool(10),
	}
}

// Requestor: Object reponsible for remote service access;
type Requestor struct {
	namingServerAddress string
	requestHandler      *client.ClientRequestHandler
	marshaller          *network.Marshaller
	namingPool          *infrastructure.NamingPool
}

// Invoke: Run desired method on remote server;
func (r *Requestor) Invoke(serviceName string, methodName string, args []interface{}) network.Response {

	serviceAddress := r.find(serviceName)
	if serviceAddress == "" {
		serviceAddress = r.lookup(serviceName)
		r.putServiceInNamingPool(serviceName, serviceAddress)
	}
	bytesRequestData := r.marshall(serviceName, methodName, args)
	serverResponse := r.send(serviceAddress, bytesRequestData)
	return r.unmarshall(serverResponse)
}

// lookup: Looks for specific remote server address from a service name;
func (r *Requestor) lookup(serviceName string) string {
	lookupResponse := r.requestHandler.Lookup(r.namingServerAddress, serviceName)
	return r.marshaller.UnmarshallLookupResponse(lookupResponse)
}

// send: Send invocation request to Client Request Handler;
func (r *Requestor) send(remoteServiceAddress string, requestData *bytes.Buffer) *http.Response {
	serverResponse := r.requestHandler.Send(remoteServiceAddress, requestData)
	return serverResponse
}

// marshall: Serializes an invoke request into a bytes buffer;
func (r *Requestor) marshall(serviceName string, methodName string, args []interface{}) *bytes.Buffer {
	clientRequest := network.NewClientRequest(serviceName, methodName, args)
	return r.marshaller.MarshallClientRequest(clientRequest)
}

// unmarshall: Deserializes an HTTP response;
func (r *Requestor) unmarshall(serverResponse *http.Response) network.Response {
	clientResponse := r.marshaller.UnmarshallClientResponse(serverResponse)
	return clientResponse
}

func (r *Requestor) find(serviceName string) string {
	response := r.namingPool.Find(serviceName)
	return response
}

func (r *Requestor) putServiceInNamingPool(serviceName string, serviceAddress string) {
	r.namingPool.PutServiceInNamingPool(serviceName, serviceAddress)
}
