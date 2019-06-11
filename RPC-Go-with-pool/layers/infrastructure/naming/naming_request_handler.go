package naming

import (
	"log"
	"net/http"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/network"
	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/services/naming"
)

// NewNamingServiceRequestHandler creates NamingServiceRequestHandler Entity
func NewNamingServiceRequestHandler() *NamingServiceRequestHandler {
	return &NamingServiceRequestHandler{
		namingService: *naming.NewNamingService(),
		marshaller:    network.NewMarshaller(),
	}
}

// NamingServiceRequestHandler Entity responsible for handle Remote Services Registration
type NamingServiceRequestHandler struct {
	marshaller    *network.Marshaller
	namingService naming.NamingService
}

// HandleLookupServices handles client lookup requests to the Naming Service for a service address
func (n *NamingServiceRequestHandler) HandleLookupServices(writer http.ResponseWriter, request *http.Request) {
	serviceName := request.URL.EscapedPath()[len("/lookup/"):]
	log.Printf("Lookup for service: '%s'.", serviceName)
	addressBytes := n.namingService.LookupService(serviceName)
	writer.Header().Set("Content-Type", "service/json; charset=utf-8")
	writer.Write(addressBytes)
}

// HandleRegistrationServices handles remote server requests to Naming Server to execute a Service Registration
func (r *NamingServiceRequestHandler) HandleRegistrationServices(writer http.ResponseWriter, request *http.Request) {
	messagesRegistrationBytes := r.namingService.RegisterServices(request)
	writer.Header().Set("Content-Type", "service/json; charset=utf-8")
	writer.Write(messagesRegistrationBytes)
}
