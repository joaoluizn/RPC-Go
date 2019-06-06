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
	DeleteObject       = "Delete"

	namingServerHost = "0.0.0.0"
	namingServerPort = "8923"
)

// type OperationArguments struct {
// 	OperationName string
// 	ProductName   string
// 	ProductPrice  float32
// }

// func NewOperation(OperationName string) *OperationArguments{
// 	return &OperationArguments{
// 		OperationName : OperationName,
// 	}
// }

// func NewOperation(OperationName string, ProductName string) *OperationArguments{
// 	return &OperationArguments{
// 		OperationName : OperationName,
// 		ProductName : ProductName,
// 	}
// }

// func NewOperation(OperationName string, ProductName string, ProductPrice float32) *OperationArguments {
// 	return &OperationArguments{
// 		OperationName: OperationName,
// 		ProductName:   ProductName,
// 		ProductPrice:  ProductPrice,
// 	}
// }

func main() {

	//Naming Server Address
	namingServerAddr := fmt.Sprintf("%s:%s", namingServerHost, namingServerPort)
	// Creating ClientProxy to call Remote Procedures
	storageClient := client.NewClientProxy(namingServerAddr, StorageServiceName)

	// namingProxyOn := true
	// storageClient.setNamingProxy(namingproxyIsOn)

	// numOfOps := 10
	// clientOperations := make([]string, numOfOps)
	// clientOperations[0] = CreateObject
	// clientOperations[1] = CreateObject
	// clientOperations[2] = CreateObject
	// clientOperations[3] = CreateObject
	// clientOperations[4] = CreateObject
	// clientOperations[5] = CreateObject
	// clientOperations[6] = CreateObject
	// clientOperations[7] = CreateObject
	// clientOperations[8] = CreateObject
	// clientOperations[9] = CreateObject

	numOfOps := 10
	clientOperations := make([]client.OperationArguments, numOfOps)
	clientOperations[0] = client.NewOperation(CreateObject, "Banana", 2.50)
	clientOperations[1] = client.NewOperation(CreateObject, "PineApple", 4.0)
	clientOperations[2] = client.NewOperation(CreateObject, "Red Apple", 1.20)
	clientOperations[3] = client.NewOperation(CreateObject, "Green Apple", 1.40)
	clientOperations[4] = client.NewOperation(CreateObject, "Watermelon", 6.30)
	clientOperations[5] = client.NewOperation(CreateObject, "Detergent", 7.50)
	clientOperations[6] = client.NewOperation(CreateObject, "Soap", 1.30)
	clientOperations[7] = client.NewOperation(CreateObject, "Shampoo", 20.00)
	clientOperations[8] = client.NewOperation(CreateObject, "Ice Cream", 23.50)
	clientOperations[9] = client.NewOperation(CreateObject, "Pizza", 19.50)

	// for i := 0; i < numOfOps; i++ {
	// 	clientOperations[i] = Operation{CreateObject, i}
	// }

	storageClient.UseRemoteService(numOfOps, clientOperations)

	time.Sleep(time.Second * 1000)

}
