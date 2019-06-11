package network

// NewService Create Service Entity
func NewService(serviceName string, serviceAddr string) *Service {
	return &Service{
		Name:    serviceName,
		Address: serviceAddr,
	}
}

// Service Entity that wraps service information
type Service struct {
	Name    string
	Address string
}

// MakeServiceList Compose a service list
func MakeServiceList(servicesNames []string, serverAddr string) []*Service {
	serviceSlice := make([]*Service, 0)

	for index := range servicesNames {
		service := NewService(servicesNames[index], serverAddr)
		serviceSlice = append(serviceSlice, service)
	}
	return serviceSlice

}
