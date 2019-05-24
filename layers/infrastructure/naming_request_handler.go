package infrastructure

import (
	"log"
	"net/http"

	"github.com/joaoluizn/RPC-go/network"
)

// NewNamingServiceRequestHandler builds a new NamingServiceRequestHandler
func NewNamingServiceRequestHandler() *NamingServiceRequestHandler {
	return &NamingServiceRequestHandler{
		// namingService: *naming.NewNamingService(),
		marshaller: network.NewMarshaller(),
	}
}

// NamingServiceRequestHandler is responsible for handle remote service's registration requests
type NamingServiceRequestHandler struct {
	marshaller *network.Marshaller
	// namingService services.NamingService
}

// HandleLookupServices handles client's look-up requests for available remote services
func (n *NamingServiceRequestHandler) HandleLookupServices(writer http.ResponseWriter, request *http.Request) {
	log.Printf("teste")
	// log.Printf("Receiving Lookup Request ", serviceName)

	serviceName := request.URL.EscapedPath()[len("/lookup/"):]
	// addressBytes := n.namingService.LookupService(serviceName)
	writer.Header().Set("Content-Type", "service/json; charset=utf-8")
	writer.Write(n.marshaller.MarshallLookupResponse(serviceName))
}

// func (n *NamingServiceRequestHandler) HandleRegisterServices(connection net.Conn, writer http.ResponseWriter, request *http.Request) {
// 	log.Printf("Receiving Register Request from ", request.RemoteAddr)
// 	// n.namingService.RegisterServices(connection)
// 	log.Printf("Cheguei ate aqui!")
// }

// // HandleRegistrationServices handles remote services registration requests
// func (n *NamingServiceRequestHandler) HandleRegistrationServices(dados string) {
// 	log.Printf("register: " + dados)
// 	// n.namingService.RegisterServices(connection)
// }
