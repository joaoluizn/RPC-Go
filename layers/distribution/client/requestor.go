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
	serviceAddress := r.lookup(serviceName)
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

// ##########################################################################################

// #########################################################

// package client

// import (
// 	"bytes"
// 	"math/rand"
// 	"net/http"
// 	"time"

// 	"github.com/joaoluizn/RPC-go/layers/infrastructure/client"
// 	"github.com/joaoluizn/RPC-go/network"
// )

// // NewRequestor: Create a Requestor instance;
// func NewRequestor(namingServerAddress string) *Requestor {
// 	return &Requestor{
// 		namingServerAddress: namingServerAddress,
// 		requestHandler:      client.NewClientRequestHandler(),
// 		marshaller:          network.NewMarshaller(),
// 		requests:            make(chan Request, 10),
// 		responses:           make(chan *http.Response),
// 	}
// }

// func (r *Requestor) NewRequest(serviceAddress string, bytesRequestData *bytes.Buffer) Request {
// 	return Request{
// 		serviceAddress:   serviceAddress,
// 		bytesRequestData: bytesRequestData,
// 	}
// }

// // Requestor: Object reponsible for remote service access;
// type Requestor struct {
// 	namingServerAddress string
// 	requestHandler      *client.ClientRequestHandler
// 	marshaller          *network.Marshaller
// 	requests            chan Request
// 	responses           chan *http.Response
// }

// type Request struct {
// 	serviceAddress   string
// 	bytesRequestData *bytes.Buffer
// }

// var (
// 	index = 0
// )

// // Invoke: Run desired method on remote server;
// func (r *Requestor) Invoke(serviceName string, methodName string, args []interface{}) network.Response {
// 	serviceAddress := r.lookup(serviceName)
// 	bytesRequestData := r.marshall(serviceName, methodName, args)
// 	r.requests <- r.NewRequest(serviceAddress, bytesRequestData)
// 	r.send(r.requests, r.responses)
// 	// r.send(r.requests, r.responses)
// 	return r.unmarshall(<-r.responses) //faz send ser uma goroutine send(request <-chan Request,esponse chan<- *httpResponse)
// }

// // lookup: Looks for specific remote server address from a service name;
// func (r *Requestor) lookup(serviceName string) string {
// 	lookupResponse := r.requestHandler.Lookup(r.namingServerAddress, serviceName)
// 	return r.marshaller.UnmarshallLookupResponse(lookupResponse)
// }

// // send: Send invocation request to Client Request Handler;
// func (r *Requestor) send(requests <-chan Request, responses chan<- *http.Response) {
// 	req := <-requests
// 	if index == 0 {
// 		time.Sleep(time.Millisecond * 500)
// 		index++
// 	}
// 	responses <- r.requestHandler.Send(req.serviceAddress, req.bytesRequestData)
// }

// // marshall: Serializes an invoke request into a bytes buffer;
// func (r *Requestor) marshall(serviceName string, methodName string, args []interface{}) *bytes.Buffer {
// 	clientRequest := network.NewClientRequest(serviceName, methodName, args)
// 	return r.marshaller.MarshallClientRequest(clientRequest)
// }

// // unmarshall: Deserializes an HTTP response;
// func (r *Requestor) unmarshall(serverResponse *http.Response) network.Response {
// 	clientResponse := r.marshaller.UnmarshallClientResponse(serverResponse)
// 	return clientResponse
// }

// ################################################

// #########################################################

// package client

// import (
// 	"bytes"
// 	"net/http"

// 	"github.com/joaoluizn/RPC-Go/layers/infrastructure"
// 	"github.com/joaoluizn/RPC-go/layers/infrastructure/client"
// 	"github.com/joaoluizn/RPC-go/network"
// )

// // NewRequestor: Create a Requestor instance;
// func NewRequestor(namingServerAddress string) *Requestor {
// 	return &Requestor{
// 		namingServerAddress: namingServerAddress,
// 		requestHandler:      client.NewClientRequestHandler(),
// 		marshaller:          network.NewMarshaller(),
// 		namingProxy:         infrastructure.NewNamingProxy(10),
// 	}
// }

// // Requestor: Object reponsible for remote service access;
// type Requestor struct {
// 	namingServerAddress string
// 	requestHandler      *client.ClientRequestHandler
// 	marshaller          *network.Marshaller
// 	namingProxy         *infrastructure.NamingProxy
// }

// // Invoke: Run desired method on remote server;
// func (r *Requestor) Invoke(serviceName string, methodName string, args []interface{}) network.Response {

// 	serviceAddress := r.find(serviceName)
// 	if serviceAddress == "" {
// 		serviceAddress = r.lookup(serviceName)
// 		r.putServiceInNamingProxy(serviceName, serviceAddress)
// 	}
// 	bytesRequestData := r.marshall(serviceName, methodName, args)
// 	serverResponse := r.send(serviceAddress, bytesRequestData)
// 	return r.unmarshall(serverResponse)
// }

// // lookup: Looks for specific remote server address from a service name;
// func (r *Requestor) lookup(serviceName string) string {
// 	lookupResponse := r.requestHandler.Lookup(r.namingServerAddress, serviceName)
// 	return r.marshaller.UnmarshallLookupResponse(lookupResponse)
// }

// // send: Send invocation request to Client Request Handler;
// func (r *Requestor) send(remoteServiceAddress string, requestData *bytes.Buffer) *http.Response {
// 	serverResponse := r.requestHandler.Send(remoteServiceAddress, requestData)
// 	return serverResponse
// }

// // marshall: Serializes an invoke request into a bytes buffer;
// func (r *Requestor) marshall(serviceName string, methodName string, args []interface{}) *bytes.Buffer {
// 	clientRequest := network.NewClientRequest(serviceName, methodName, args)
// 	return r.marshaller.MarshallClientRequest(clientRequest)
// }

// // unmarshall: Deserializes an HTTP response;
// func (r *Requestor) unmarshall(serverResponse *http.Response) network.Response {
// 	clientResponse := r.marshaller.UnmarshallClientResponse(serverResponse)
// 	return clientResponse
// }

// func (r *Requestor) find(serviceName string) string {
// 	response := r.namingProxy.Find(serviceName)
// 	return response
// }

// func (r *Requestor) putServiceInNamingProxy(serviceName string, serviceAddress string) {
// 	r.namingProxy.PutServiceInNamingProxy(serviceName, serviceAddress)
// }
