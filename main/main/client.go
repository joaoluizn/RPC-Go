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

type Operation struct {
	operationName string
	operationId   int
}

type Response struct {
	operationResponse interface{}
	operationId       int
}

func useRemoteService(client *client.ClientProxy, numOfOps int, clientOperations []Operation) {

	operations := make(chan Operation, numOfOps)
	responses := make(chan Response, numOfOps)

	go UseRemoteService(client, operations, responses)
	go UseRemoteService(client, operations, responses)
	go UseRemoteService(client, operations, responses)
	go UseRemoteService(client, operations, responses)
	go UseRemoteService(client, operations, responses)

	for i := 0; i < numOfOps; i++ {
		operations <- clientOperations[i]
	}
	close(operations)

	for r := 0; r < numOfOps; r++ {
		response := <-responses
		log.Printf("Response of Operation %d: %s", response.operationId, response.operationResponse)
	}
}

func UseRemoteService(client *client.ClientProxy, operations <-chan Operation, responses chan<- Response) {

	time.Sleep(time.Second * 2)

	for op := range operations {
		// Calling a Remote Procedure
		log.Printf("Operation %d: Calling Remote Procedure: '%s'", op.operationId, op.operationName)
		// This Invoke can receive the operation to be executed and arguments needed
		responses <- Response{(client.Invoke(op.operationName, "teste", op.operationId).Content[0]), op.operationId}
	}
}

func main() {

	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)
	// Creating ClientProxy to call Remote Procedures
	storageClient := client.NewClientProxy(namingServerAddr, StorageServiceName)

	// namingProxyOn := true
	// storageClient.setNamingProxy(namingproxyIsOn)

	numOfOps := 10
	clientOperations := make([]Operation, numOfOps)
	clientOperations[0] = Operation{CreateObject, 0}
	clientOperations[1] = Operation{CreateObject, 1}
	clientOperations[2] = Operation{CreateObject, 2}
	clientOperations[3] = Operation{CreateObject, 3}
	clientOperations[4] = Operation{CreateObject, 4}
	clientOperations[5] = Operation{CreateObject, 5}
	clientOperations[6] = Operation{CreateObject, 6}
	clientOperations[7] = Operation{CreateObject, 7}
	clientOperations[8] = Operation{CreateObject, 8}
	clientOperations[9] = Operation{CreateObject, 9}

	// for i := 0; i < numOfOps; i++ {
	// 	clientOperations[i] = Operation{CreateObject, i}
	// }

	useRemoteService(storageClient, numOfOps, clientOperations)

	time.Sleep(time.Second * 1000)

}
