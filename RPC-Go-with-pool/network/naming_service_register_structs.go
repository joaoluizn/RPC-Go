package network

// NewNamingServiceRegistration create NamingServiceRegistration Entity;
func NewNamingServiceRegistration(serviceName []string, serverAddress string) *NamingServiceRegistration {
	return &NamingServiceRegistration{
		ServicesNames: serviceName, ServerAddress: serverAddress,
	}
}

// NamingServiceRegistration Entity responsible to wrap a services and address for Registration puposes on Naming Service;
type NamingServiceRegistration struct {
	ServicesNames []string
	ServerAddress string
}
