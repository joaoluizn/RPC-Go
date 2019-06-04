package main

import (
	"fmt"

	"github.com/joaoluizn/RPC-go/layers/infrastructure/storage"
)

func main() {
	storageServerHost := "0.0.0.0"

	storageServerPort_0 := "8925"
	// storageServerPort_1 := "8926"
	// storageServerPort_2 := "8927"

	namingServerHost := "0.0.0.0"
	namingServerPort := "8923"

	storageServerAddr_0 := fmt.Sprintf("%s:%s", storageServerHost, storageServerPort_0)
	// storageServerAddr_1 := fmt.Sprintf("%s:%s", storageServerHost, storageServerPort_1)
	// storageServerAddr_2 := fmt.Sprintf("%s:%s", storageServerHost, storageServerPort_2)

	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)

	remoteServiceServer_0 := storage.NewStorageServiceServer(storageServerAddr_0, namingServerAddr)
	// remoteServiceServer_1 := storage.NewStorageServiceServer(storageServerAddr_1, namingServerAddr)
	// remoteServiceServer_2 := storage.NewStorageServiceServer(storageServerAddr_2, namingServerAddr)

	// There could be many services, so one register line for each service in that remote IP
	remoteServiceServer_0.RegisterServiceInLocalStorage(storage.StorageServiceName, storage.NewStorage())
	// remoteServiceServer_1.RegisterServiceInLocalStorage(storage.StorageServiceName, storage.NewStorage())
	// remoteServiceServer_2.RegisterServiceInLocalStorage(storage.StorageServiceName, storage.NewStorage())

	remoteServiceServer_0.RunRemoteServer()
	// remoteServiceServer_1.RunRemoteServer()
	// remoteServiceServer_2.RunRemoteServer()

}
