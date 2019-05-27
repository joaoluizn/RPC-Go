package storage

import (
	"net/http"

	"github.com/joaoluizn/RPC-go/layers/distribution/server"
)

// NewRemoteServiceRequestHandler builds a new RemoteServiceRequestHandler
func NewStorageServiceRequestHandler() *StorageServiceRequestHandler {
	return &StorageServiceRequestHandler{
		Invoker: server.NewInvoker(),
	}
}

// RemoteServiceRequestHandler is responsible for handle client's invocation requests
type StorageServiceRequestHandler struct {
	Invoker *server.Invoker
}

//HandleInvokeRequest handles client's requests
func (r *StorageServiceRequestHandler) HandleInvokeRequest(writer http.ResponseWriter, request *http.Request) {
	// log.Printf(internal.MsgClientInvokeRequest, request.RemoteAddr)
	output := r.Invoker.Invoke(request)
	writer.Header().Set("Content-Type", "service/json; charset=utf-8")
	writer.Write(output)
}
