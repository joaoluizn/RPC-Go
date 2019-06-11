package naming

import (
	"log"
	"net"
	"net/http"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/network"
)

// NewNamingServiceServer create a new instance of NamingServiceServer
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

// RunNamingServer core function to run the naming service
func (n *NamingServiceServer) RunNamingServer() {
	go n.runHTTPServerForServiceLookup()
	n.runHTTPForServicesRegistration()
}

// runHTTPServerForServiceLookup Run a http server for the /lookup endpoint
func (n *NamingServiceServer) runHTTPServerForServiceLookup() {
	defer n.listener.Close()
	log.Printf("Running Lookup Endpoint on %s\n\n", n.namnigServerAddr)

	http.HandleFunc("/lookup/", n.requestHandler.HandleLookupServices)

	serveErr := http.Serve(n.listener, nil)
	if serveErr != nil {
		log.Fatal(serveErr.Error())
	}
}

// runHTTPForServicesRegistration Run a http server for the /register endpoint
func (n *NamingServiceServer) runHTTPForServicesRegistration() {
	defer n.listener.Close()

	log.Printf("Running Services Registration on %s", n.namnigServerAddr)
	for {
		http.HandleFunc("/register/", n.requestHandler.HandleRegistrationServices)

		serveErr := http.Serve(n.listener, nil)
		if serveErr != nil {
			log.Fatal(serveErr.Error())
		}
	}
}
