package main

import (
	"fmt"

	"github.com/joaoluizn/RPC-go/layers/infrastructure/naming"
)

func main() {
	namingServerHost := "0.0.0.0"
	namingServerPort := "8923"

	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)

	namingServiceServer := naming.NewNamingServiceServer(namingServerAddr)
	namingServiceServer.Run()
}