package main

import (
	"fmt"
	"time"
	"strconv"

	"github.com/joaoluizn/RPC-Go/RPC-Go-simple/layers/distribution/client"
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

	numOfOps := 10000
	clientOperations := make([]client.OperationArguments, 10)
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

	for i := 0; i < numOfOps; i++ {
		start := time.Now()
		storageClient.UseRemoteService(clientOperations[i%10])
		elapsed := time.Since(start)
		fmt.Printf(strconv.FormatInt(elapsed.Nanoseconds(),10)+"\n")
	}

}
