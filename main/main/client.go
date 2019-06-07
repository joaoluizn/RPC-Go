package main

import (
	"fmt"
	"time"

	"github.com/joaoluizn/RPC-Go/layers/distribution/client"
)

const (
	StorageServiceName = "Storage"
	HelloStorage       = "HelloStorage"
	CreateObject       = "Create"
	RemoveObject       = "Remove"
	UpdateObject       = "Update"
	ReadObjectList     = "ReadList"
	DeleteObject       = "Delete"

	namingServerHost = "0.0.0.0"
	namingServerPort = "8923"
)

func main() {

	//Naming Server Address
	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)
	// Creating ClientProxy to call Remote Procedures
	storageClient := client.NewClientProxy(namingServerAddr, StorageServiceName)

	numOfOps := 10
	clientOperations := make([]client.OperationArguments, numOfOps)
	clientOperations[0] = client.NewOperation(CreateObject, "Banana", 2.95)
	clientOperations[1] = client.NewOperation(CreateObject, "PineApple", 4.95)
	clientOperations[2] = client.NewOperation(CreateObject, "Red Apple", 1.25)
	clientOperations[3] = client.NewOperation(CreateObject, "Green Apple", 1.75)
	clientOperations[4] = client.NewOperation(CreateObject, "Watermelon", 6.25)
	clientOperations[5] = client.NewOperation(CreateObject, "Detergent", 7.25)
	clientOperations[6] = client.NewOperation(CreateObject, "Soap", 1.95)
	clientOperations[7] = client.NewOperation(CreateObject, "Shampoo", 20.25)
	clientOperations[8] = client.NewOperation(CreateObject, "Ice Cream", 23.95)
	clientOperations[9] = client.NewOperation(CreateObject, "Pizza", 19.95)

	storageClient.UseRemoteService(numOfOps, clientOperations)

	time.Sleep(time.Second * 1)

	readOperation := make([]client.OperationArguments, 1)
	readOperation[0] = client.NewOperation(ReadObjectList, "", 0)

	storageClient.UseRemoteService(1, readOperation)

	time.Sleep(time.Second * 1)

	updateOperation := make([]client.OperationArguments, 1)
	updateOperation[0] = client.NewOperation(UpdateObject, "PineApple", 5000.95)

	time.Sleep(time.Second * 1)

	storageClient.UseRemoteService(1, updateOperation)

	time.Sleep(time.Second * 1)

	storageClient.UseRemoteService(1, readOperation)

	time.Sleep(time.Second * 1000)
}
