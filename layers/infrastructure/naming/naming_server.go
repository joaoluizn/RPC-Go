package naming

import (
	"log"
	"net"
	"net/http"

	"github.com/joaoluizn/RPC-go/network"
)

// NewNamingServiceServer builds a new instance of NamingServiceServer
func NewNamingServiceServer(namingServerAddr string) *NamingServiceServer {
	return &NamingServiceServer{
		requestHandler:   NewNamingServiceRequestHandler(),
		listener:         network.GetTCPListener(namingServerAddr),
		namnigServerAddr: namingServerAddr,
	}
}

// NamingServiceServer handles the address for services available for clients
type NamingServiceServer struct {
	requestHandler   *NamingServiceRequestHandler
	listener         *net.TCPListener
	namnigServerAddr string
}

// Run runs the naming service
func (n *NamingServiceServer) Run() {
	go n.runHTTPServerForServiceLookup()
	n.runSocketForServicesRegistration()
}

// runHTTPServerForServiceLookup runs a http server for remote services look-up
func (n *NamingServiceServer) runHTTPServerForServiceLookup() {
	defer n.listener.Close()

	// Mapping lookup route to NameRequestHandler
	http.HandleFunc("/lookup/", n.requestHandler.HandleLookupServices)

	serve_err := http.Serve(n.listener, nil)
	if serve_err != nil {
		log.Fatal(serve_err.Error())
	}
}

// runSocketForServicesRegistration runs a network socket for remote services registration
func (n *NamingServiceServer) runSocketForServicesRegistration() {
	defer n.listener.Close()

	log.Printf("Running Services Registration on %s", n.namnigServerAddr)
	for {
		http.HandleFunc("/register/", n.requestHandler.HandleRegistrationServices)

		serve_err := http.Serve(n.listener, nil)
		if serve_err != nil {
			log.Fatal(serve_err.Error())
		}
	}
}