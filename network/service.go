package network

func NewService(serviceName string, serviceAddr string) *Service {
	return &Service{
		Name:    serviceName,
		Address: serviceAddr,
	}
}

type Service struct {
	Name    string
	Address string
}

func MakeServiceList(servicesNames []string, serverAddr string) []*Service {
	// Create service slice
	serviceSlice := make([]*Service, 0)

	for index := range servicesNames {
		service := NewService(servicesNames[index], serverAddr)
		serviceSlice = append(serviceSlice, service)
	}
	return serviceSlice

}

