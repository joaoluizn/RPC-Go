package naming

import (
	"log"
	"net/http"

	"github.com/joaoluizn/RPC-go/infrastructure"
	"github.com/joaoluizn/RPC-go/network"
)

// NewNamingServiceServer builds a new instance of NamingServiceServer
func NewNamingServiceServer(namingServerPort string) *NamingServiceServer {
	return &NamingServiceServer{
		requestHandler:   infrastructure.NewNamingServiceRequestHandler(),
		namingServerPort: namingServerPort,
	}
}

// NamingServiceServer handles the address for services available for clients
type NamingServiceServer struct {
	requestHandler   *infrastructure.NamingServiceRequestHandler
	namingServerPort string
}

// Run runs the naming service
func (n *NamingServiceServer) Run() {
	go n.runHTTPServerForServiceLookup()
	go n.runSocketForServicesRegistration()
	log.Printf("Running Naming and Register Service.")
}

// runHTTPServerForServiceLookup runs a http server for remote services look-up
func (n *NamingServiceServer) runHTTPServerForServiceLookup() {
	listener, _ := network.GetTCPListener(n.namingServerPort)

	// Mapping lookup route to NameRequestHandler
	http.HandleFunc("/lookup/", n.requestHandler.HandleLookupServices)

	serve_err := http.Serve(listener, nil)
	if serve_err != nil {
		log.Fatal(serve_err.Error())
	}
}

// runSocketForServicesRegistration runs a network socket for remote services registration
func (n *NamingServiceServer) runSocketForServicesRegistration() {
	listener, _ := network.GetTCPListener(n.namingServerPort)
	defer listener.Close()
	// log.Printf(internal.MsgRunningServicesRegistration, address)

	connection, err := listener.Accept()
	if err != nil {
		log.Println(err.Error())
	}

	http.HandleFunc("/register/", func(output http.ResponseWriter, request *http.Request) {
		n.requestHandler.HandleRegisterServices(connection, output, request)
	})

	serve_err := http.Serve(listener, nil)
	if serve_err != nil {
		log.Fatal(serve_err.Error())
	}

	// for {
	// 	connection, err := listener.Accept()
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 	}

	// 	go n.requestHandler.HandleRegistrationServices(connection)
	// }
}
