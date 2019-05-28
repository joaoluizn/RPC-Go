package main

import (
	"fmt"

	"github.com/joaoluizn/RPC-go/layers/infrastructure/naming"
)

const (
	namingServerHost = "0.0.0.0"
	namingServerPort = "8923"
)

func main() {
	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)

	namingServiceServer := naming.NewNamingServiceServer(namingServerAddr)
	namingServiceServer.RunNamingServer()
}
