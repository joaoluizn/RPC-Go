package network

import (
	"log"
	"net"
)

// GetTCPListener gets a TCP listener choosing the first available port
func GetTCPListener(serviceAddr string) *net.TCPListener {
	address, _ := net.ResolveTCPAddr("tcp", serviceAddr)
	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		log.Fatal("utils (GetTCPListener): ", err.Error())
	}
	return listener
}

// GetTCPDialer gets a TCP dialer to communicate with a TCP socket
func GetTCPDialer(address string) *net.TCPConn {
	tcpAddress, _ := net.ResolveTCPAddr("tcp", address)
	dialer, _ := net.DialTCP("tcp", nil, tcpAddress)
	return dialer

}
