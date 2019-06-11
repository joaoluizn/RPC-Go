package infrastructure

// import "log"

// NewConnectionPool:  Create ConnectionPool instance;
func NewNamingProxy(size int) *NamingProxy {
	return &NamingProxy{
		ServiceBuffer: make(map[string]string, size),
		PoolSize:      size,
		On:            false,
	}
}

// Naming:
type NamingProxy struct {
	ServiceBuffer map[string]string
	PoolSize      int
	On            bool
}

func (p *NamingProxy) Find(serviceName string) string {

	_, nameInNamingProxy := p.ServiceBuffer[serviceName]

	if nameInNamingProxy {
		// log.Printf("Service Found in Naming Proxy ")
		return p.ServiceBuffer[serviceName]
	}

	// log.Print("Service Not Found in Naming Proxy.")
	// log.Printf("Naming Server Address: ")
	return ""

}

func (p *NamingProxy) PutServiceInNamingProxy(serviceName string, serviceAddress string) {
	p.ServiceBuffer[serviceName] = serviceAddress
}
