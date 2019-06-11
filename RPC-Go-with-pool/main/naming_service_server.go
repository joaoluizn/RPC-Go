package main

import (
	"fmt"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/layers/infrastructure/naming"
)

const (
	namingServerHost = "0.0.0.0"
	namingServerPort = "8923"
)

// main Naming Server Main function
func main() {
	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)

	namingServiceServer := naming.NewNamingServiceServer(namingServerAddr)
	namingServiceServer.RunNamingServer()
}
