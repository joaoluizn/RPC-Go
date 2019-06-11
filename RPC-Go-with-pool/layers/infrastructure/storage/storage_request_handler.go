package storage

import (
	"net/http"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/layers/distribution/server"
)

// NewStorageServiceRequestHandler create RemoteServiceRequestHandler Entity
func NewStorageServiceRequestHandler() *StorageServiceRequestHandler {
	return &StorageServiceRequestHandler{
		Invoker: server.NewInvoker(),
	}
}

// StorageServiceRequestHandler Entity responsible for Handle Invocation Requests from Client
type StorageServiceRequestHandler struct {
	Invoker *server.Invoker
}

//HandleInvokeRequest Handle Invocation Requests from Client
func (r *StorageServiceRequestHandler) HandleInvokeRequest(writer http.ResponseWriter, request *http.Request) {
	output := r.Invoker.Invoke(request)
	writer.Header().Set("Content-Type", "service/json; charset=utf-8")
	writer.Write(output)
}
