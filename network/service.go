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
	serviceList := make([]*Service, 0)

	for index := range servicesNames {
		service := NewService(servicesNames[index], serverAddr)
		serviceList = append(serviceList, service)
	}
	return serviceList
}
