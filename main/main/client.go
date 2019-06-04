package main

import (
	"fmt"
	"log"
	"time"

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

var (
	id = 0
)

func UseRemoteService(client *client.ClientProxy, methodName string, test int, ID int) {
	// Calling a Remote Procedure
	log.Printf("Operation %d: Calling Remote Procedure: '%s'", ID, methodName)
	// This Invoke can receive the operation to be executed and arguments needed
	helloResponse := client.Invoke(methodName, "teste", test)
	// createResponse
	log.Printf("Response of Operation %d: %s", ID, helloResponse.Content[0])
	// fmt.Printf("\n")
	id++
}

func main() {

	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)
	// namingServerAddr := namingServerHost + ":" + namingServerPort

	// Creating ClientProxy to call Remote Procedures
	storageClient0 := client.NewClientProxy(namingServerAddr, StorageServiceName)

	for i := 0; i < 4; i++ {
		go UseRemoteService(storageClient0, CreateObject, 20, i)
	}

	// go UseRemoteService(storageClient0, Create
	time.Sleep(time.Millisecond * 5000)

	//nao aceita nenhuma 5ª operacao
	// storageClient1 := client.NewClientProxy(namingServerAddr, StorageServiceName)

}
