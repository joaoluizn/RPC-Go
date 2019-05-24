package main

import "github.com/joaoluizn/RPC-go/layers/infrastructure"

func main() {
	namingServerAddr := "0.0.0.0"
	namingServerPort := "8923"

	remoteServiceServer := infrastructure.NewStorageServiceServer(namingServerAddr, namingServerPort)
	remoteServiceServer.RegisterServiceInNamingService(ServiceName, NewStorage())
	remoteServiceServer.Run()
}
