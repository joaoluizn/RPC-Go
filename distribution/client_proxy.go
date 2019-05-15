package distribution

import (
	"../network"
)

// NewClientProxy: Create ClientProxy instance;
func NewClientProxy(namingServerAddress string, serviceName string) *ClientProxy {
	return &ClientProxy{
		requestor:   NewRequestor(namingServerAddress),
		serviceName: serviceName,
	}
}

// ClientProxy: Object reponsible for remote communication
type ClientProxy struct {
	requestor   *Requestor
	serviceName string
}

// Invoke: Run desired method on remote server;
func (p *ClientProxy) Invoke(methodName string, arguments ...interface{}) network.Response {
	return p.requestor.Invoke(p.serviceName, methodName, arguments)
}
