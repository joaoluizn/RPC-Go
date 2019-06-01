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
	log.Printf("Calling Remote Procedure: '%s'", CreateObject)

	// This Invoke can receive the operation to be executed and arguments needed
	helloResponse := storageClient.Invoke(CreateObject, "teste", 20)

	// createResponse := storageClient.Invoke(CreateStorage, args)
	log.Printf("Response: %s", helloResponse.Content[0])

	log.Printf("Calling Remote Procedure: '%s'", CreateObject)
	helloResponse_2 := storageClient.Invoke(CreateObject, "teste", 20)
	log.Printf("Response: %s", helloResponse_2.Content[0])
}
