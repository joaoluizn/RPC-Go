package network

import (
	"log"
	"net"
)

// GetTCPListener gets a TCP listener choosing the first available port
func GetTCPListener(addr string, port string) *net.TCPListener {
	log.Printf(addr + ":" + port)
	address, _ := net.ResolveTCPAddr("tcp", addr+":"+port)
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

// buildAddress builds a network address compose by [host_ip]:[port]
// func buildAddress(port string) string {
// 	return getLocalIP() + ":" + port
// }

// // getLocalIP gets the host ip address
// func getLocalIP() string {
// 	addresses, err := net.InterfaceAddrs()
// 	var hostIP string
// 	if err != nil {
// 		hostIP = ""
// 	}

// 	index := 0
// 	for hostIP == "" && index < len(addresses) {
// 		address := addresses[index]
// 		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
// 			if ipnet.IP.To4() != nil {
// 				hostIP = ipnet.IP.String()
// 			}
// 		}
// 		index++
// 	}
// 	return hostIP
// }
