package network

import (
	"net"
)

func NewService(connection net.Conn, serviceName string, serviceAddr string) *Service {
	return &Service{
		Connection: connection,
		Name:       serviceName,
		Address:    serviceAddr,
	}
}

type Service struct {
	Connection net.Conn
	Name       string
	Address    string
}

// func NewRemoteServices(connection net.Conn, servicesNames []string, serverAddr string) []*Service {
// 	services_list := make([]*Service, 0)

// 	for index := range servicesNames {
// 		service := NewRemoteServiceEntry(connection, servicesNames[index], serverAddress)
// 		services = append(services_list, service)
// 	}
// 	return services
// }
