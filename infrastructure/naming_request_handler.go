package infrastructure

import (
	"log"
	"net"
	"net/http"
)

// NewNamingServiceRequestHandler builds a new NamingServiceRequestHandler
func NewNamingServiceRequestHandler() *NamingServiceRequestHandler {
	return &NamingServiceRequestHandler{
		// namingService: *naming.NewNamingService(),
	}
}

// NamingServiceRequestHandler is responsible for handle remote service's registration requests
type NamingServiceRequestHandler struct {
	// namingService services.NamingService
}

// HandleLookupServices handles client's look-up requests for available remote services
func (n *NamingServiceRequestHandler) HandleLookupServices(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Receiving Lookup Request from ", request.RemoteAddr)
	// serviceName := request.URL.EscapedPath()[len("/lookup/"):]
	// addressBytes := n.namingService.LookupService(serviceName)
	// writer.Header().Set(network.HeaderContentTypeTag, network.HeaderApplicationJSONUTF8)
	// writer.Write(addressBytes)
}

func (n *NamingServiceRequestHandler) HandleRegisterServices(connection net.Conn, writer http.ResponseWriter, request *http.Request) {
	log.Printf("Receiving Register Request from ", request.RemoteAddr)
	// n.namingService.RegisterServices(connection)
	log.Printf("Cheguei ate aqui!")
}

// HandleRegistrationServices handles remote services registration requests
// func (n *NamingServiceRequestHandler) HandleRegistrationServices(connection net.Conn) {
// 	n.namingService.RegisterServices(connection)
// }
