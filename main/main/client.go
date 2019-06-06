package main

import (
	"fmt"
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

func main() {

	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)
	// Creating ClientProxy to call Remote Procedures
	storageClient := client.NewClientProxy(namingServerAddr, StorageServiceName)

	// namingProxyOn := true
	// storageClient.setNamingProxy(namingproxyIsOn)

	numOfOps := 10
	clientOperations := make([]string, numOfOps)
	clientOperations[0] = CreateObject
	clientOperations[1] = CreateObject
	clientOperations[2] = CreateObject
	clientOperations[3] = CreateObject
	clientOperations[4] = CreateObject
	clientOperations[5] = CreateObject
	clientOperations[6] = CreateObject
	clientOperations[7] = CreateObject
	clientOperations[8] = CreateObject
	clientOperations[9] = CreateObject

	// for i := 0; i < numOfOps; i++ {
	// 	clientOperations[i] = Operation{CreateObject, i}
	// }

	storageClient.UseRemoteService(numOfOps, clientOperations)

	time.Sleep(time.Second * 1000)

}
