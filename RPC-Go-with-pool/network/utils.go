package network

import (
	"log"
	"net"
)

// GetTCPListener Get a TCP listener to listen on a specific port
func GetTCPListener(serviceAddr string) *net.TCPListener {
	address, _ := net.ResolveTCPAddr("tcp", serviceAddr)
	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		log.Fatal("utils (GetTCPListener): ", err.Error())
	}
	return listener
}

// GetTCPDialer Get a TCP dialer to communication purposes with a specific address
func GetTCPDialer(address string) *net.TCPConn {
	tcpAddress, _ := net.ResolveTCPAddr("tcp", address)
	dialer, _ := net.DialTCP("tcp", nil, tcpAddress)
	return dialer

}
