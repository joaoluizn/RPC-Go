package naming

import (
	"log"
	"net/http"

	"github.com/joaoluizn/RPC-go/network"
	"github.com/joaoluizn/RPC-go/services/naming"
)

// NewNamingServiceRequestHandler builds a new NamingServiceRequestHandler
func NewNamingServiceRequestHandler() *NamingServiceRequestHandler {
	return &NamingServiceRequestHandler{
		namingService: *naming.NewNamingService(),
		marshaller:    network.NewMarshaller(),
	}
}

// NamingServiceRequestHandler is responsible for handle remote service's registration requests
type NamingServiceRequestHandler struct {
	marshaller    *network.Marshaller
	namingService naming.NamingService
}

// HandleLookupServices handles client's look-up requests for available remote services
func (n *NamingServiceRequestHandler) HandleLookupServices(writer http.ResponseWriter, request *http.Request) {
	serviceName := request.URL.EscapedPath()[len("/lookup/"):]
	log.Printf("Lookup for service: '%s'.", serviceName)
	addressBytes := n.namingService.LookupService(serviceName)
	writer.Header().Set("Content-Type", "service/json; charset=utf-8")
	writer.Write(addressBytes)
}

func (r *NamingServiceRequestHandler) HandleRegistrationServices(writer http.ResponseWriter, request *http.Request) {
	messagesRegistrationBytes := r.namingService.RegisterServices(request)
	writer.Header().Set("Content-Type", "service/json; charset=utf-8")
	writer.Write(messagesRegistrationBytes)
}
