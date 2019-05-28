package main

import (
	"fmt"
	"log"

	"github.com/joaoluizn/RPC-go/layers/distribution/client"
)

const (
	StorageServiceName = "Storage"
	HelloStorage       = "HelloStorage"
	CreateObject       = "Create"
	RemoveObject       = "Remove"
	UpdateObject       = "Update"
	DeleteObject       = "Delete"

	namingServerHost = "0.0.0.0"
	namingServerPort = "8923"
)

func main() {

	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)

	// Creating ClientProxy to call Remote Procedures

	storageClient := client.NewClientProxy(namingServerAddr, StorageServiceName)

	// Calling a Remote Procedure
	log.Printf("Calling Remote Procedure: '%s'", HelloStorage)

	helloResponse := storageClient.Invoke(HelloStorage)
	log.Printf("Response: %s", helloResponse.Content[0])

	// This Invoke can receive the operation to be executed and arguments needed.
	// createResponse := storageClient.Invoke(CreateStorage, args)
}
