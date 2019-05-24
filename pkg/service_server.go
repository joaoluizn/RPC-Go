package pkg

import (
	"log"
	"net"
	"net/http"
)

//objeto que tem o endereço do serviço e o endereço do register
//esse objeto pode receber um invoke do cliente
type ServiceServer struct {
	requestHandler            *infrastructure.ServiceRequestHandler
	port                      string
	registrationServerAddress string
}

//cria um novo server
func NewServiceServer(port string, registrationServerAddress string) *ServiceServer {
	return &RemoteServiceServer{
		requestHandler:            infrastructure.NewServiceRequestHandler(),
		port:                      port,
		registrationServerAddress: registrationServerAddress,
	}
}

// adiciona um novo serviço que ficará disponível no serviço de nomes
func (r *ServiceServer) RegisterInNamingService(name string, instance interface{}) {
	r.requestHandler.Invoker.Service.RegisterService(name, instance)
}

// chama o servidor http  que lida com invoke do serviço
func ServerHTTPForServiceInvocation(listener net.Listener, address string) {
	log.Printf("remote service running... invoke at address \"%s\"", address)
	http.HandleFunc(network.InvokePath)
}
