package main

import "github.com/joaoluizn/RPC-go/layers/infrastructure/naming"

func main() {
	namingServerAddr := "0.0.0.0"
	namingServerPort := "8923"

	namingServiceServer := naming.NewNamingServiceServer(namingServerAddr, namingServerPort)
	namingServiceServer.Run()
}
