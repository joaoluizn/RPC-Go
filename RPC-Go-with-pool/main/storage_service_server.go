package main

import (
	"fmt"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/layers/infrastructure/storage"
)

// main Remote Server Main Function
func main() {
	storageServerHost := "0.0.0.0"

	storageServerPort := "8925"
	namingServerHost := "0.0.0.0"
	namingServerPort := "8923"

	storageServerAddr := fmt.Sprintf("%s:%s", storageServerHost, storageServerPort)
	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)

	remoteServiceServer := storage.NewStorageServiceServer(storageServerAddr, namingServerAddr)

	// There could be many services, so one register line for each service in that remote IP
	remoteServiceServer.RegisterServiceInLocalStorage(storage.StorageServiceName, storage.NewStorage())

	remoteServiceServer.RunRemoteServer()
}
