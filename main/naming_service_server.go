package main

import "github.com/joaoluizn/RPC-go/layers/infrastructure"

func main() {
	namingServerAddr := "0.0.0.0"
	namingServerPort := "8923"

	namingServiceServer := infrastructure.NewNamingServiceServer(namingServerAddr, namingServerPort)
	namingServiceServer.Run()
}
