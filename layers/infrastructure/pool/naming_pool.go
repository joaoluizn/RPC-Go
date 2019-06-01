package infrastructure

import (
	"log"
)

// NewConnectionPool:  Create ConnectionPool instance;
func NewNamingPool(size int) *NamingPool {
	return &NamingPool{
		ServiceBuffer: make(map[string]string, size),
	}
}

// Naming:
type NamingPool struct {
	ServiceBuffer map[string]string
	PoolSize      int
}

func (p *NamingPool) Find(serviceName string) string {

	_, nameInPool := p.ServiceBuffer[serviceName]

	if nameInPool {
		log.Printf("Service Found in Naming Pool ")
		return p.ServiceBuffer[serviceName]
	}

	log.Print("Service Not Found in Naming Pool.")
	// log.Printf("Naming Server Address: ")
	return ""

}

func (p *NamingPool) PutServiceInNamingPool(serviceName string, serviceAddress string) {
	p.ServiceBuffer[serviceName] = serviceAddress
}
