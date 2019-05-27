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
	log.Printf("Receiving Lookup Request ", serviceName)

	// addressBytes := n.namingService.LookupService(serviceName)
	writer.Header().Set("Content-Type", "service/json; charset=utf-8")
	writer.Write(n.marshaller.MarshallLookupResponse(serviceName))
}

// func (n *NamingServiceRequestHandler) HandleRegisterServices(connection net.Conn, writer http.ResponseWriter, request *http.Request) {
// 	log.Printf("Receiving Register Request from ", request.RemoteAddr)
// 	// n.namingService.RegisterServices(connection)
// 	log.Printf("Cheguei ate aqui!")
// }

// HandleRegistrationServices handles remote services registration requests
// func (n *NamingServiceRequestHandler) HandleRegistrationServices(registerData string) {
// 	log.Printf("register: " + registerData)
// 	// n.namingService.RegisterServices(connection)
// }

func (r *NamingServiceRequestHandler) HandleRegistrationServices(writer http.ResponseWriter, request *http.Request) {
	r.namingService.RegisterServices(request)

	// r.marshaller.UnmarshalNamingServiceRegistration(request.Body)
	// body, err := ioutil.ReadAll(request.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println(string(body))
	// writer.Header().Set("Content-Type", "service/json; charset=utf-8")
	// writer.Write(output)
}
