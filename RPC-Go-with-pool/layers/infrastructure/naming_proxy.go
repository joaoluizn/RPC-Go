package infrastructure

import "log"

// NewNamingProxy Create ConnectionPool instance;
func NewNamingProxy(size int) *NamingProxy {
	return &NamingProxy{
		ServiceBuffer: make(map[string]string, size),
		PoolSize:      size,
		On:            false,
	}
}

// NamingProxy Object to stora previously used proxy for otimization puposes
type NamingProxy struct {
	ServiceBuffer map[string]string
	PoolSize      int
	On            bool
}

// Find search a service in the Naming Proxy
func (p *NamingProxy) Find(serviceName string) string {
	_, nameInNamingProxy := p.ServiceBuffer[serviceName]

	if nameInNamingProxy {
		log.Printf("Service Found in Naming Proxy ")
		return p.ServiceBuffer[serviceName]
	}

	log.Print("Service Not Found in Naming Proxy.")
	return ""

}

// PutServiceInNamingProxy Save a service in the naming Proxy
func (p *NamingProxy) PutServiceInNamingProxy(serviceName string, serviceAddress string) {
	p.ServiceBuffer[serviceName] = serviceAddress
}
