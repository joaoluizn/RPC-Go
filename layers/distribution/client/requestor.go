package client

import (
	"bytes"
	"net/http"

	"github.com/joaoluizn/RPC-go/layers/infrastructure/client"
	"github.com/joaoluizn/RPC-go/network"
)

// NewRequestor: Create a Requestor instance;
func NewRequestor(namingServerAddress string) *Requestor {
	return &Requestor{
		namingServerAddress: namingServerAddress,
		requestHandler:      client.NewClientRequestHandler(),
		marshaller:          network.NewMarshaller(),
	}
}

// Requestor: Object reponsible for remote service access;
type Requestor struct {
	namingServerAddress string
	requestHandler      *client.ClientRequestHandler
	marshaller          *network.Marshaller
}

// Invoke: Run desired method on remote server;
func (r *Requestor) Invoke(serviceName string, methodName string, args []interface{}) network.Response {
	remoteServiceAddress := r.lookup(serviceName)
	bytesRequestData := r.marshall(serviceName, methodName, args)
	serverResponse := r.send(remoteServiceAddress, bytesRequestData)
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
