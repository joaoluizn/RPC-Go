package naming

import (
	"github.com/joaoluizn/RPC-go/network"
)

// NewNamingService builds a new instance of NamingService
func NewNamingService() *NamingService {
	return &NamingService{
		remoteServicesEntries: make(map[string][]*Service),
		marshaller:            network.NewMarshaller(),
	}
}

// NamingService is a naming service who holds the urls to all services available for client
type NamingService struct {
	remoteServicesEntries map[string][]*Service
	marshaller            *network.Marshaller
}
