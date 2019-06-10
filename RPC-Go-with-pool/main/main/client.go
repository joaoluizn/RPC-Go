package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/layers/distribution/client"
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

	numOfOps := 100
	clientOperations := make([]client.OperationArguments, numOfOps)

	for i := 0; i < numOfOps; i++ {
		clientOperations[i] = client.NewOperation(CreateObject, "Banana#"+strconv.Itoa(i), 2.95)
	}

	start := time.Now()
	storageClient.UseRemoteService(numOfOps, clientOperations)
	elapsed := time.Since(start)
	fmt.Printf(strconv.FormatInt(elapsed.Nanoseconds(), 10) + "\n")
}
