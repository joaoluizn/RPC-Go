package main

import "github.com/joaoluizn/RPC-go/layers/distribution/client"

func main() {
	storageClient := client.NewClientProxy("0.0.0.0:8923", "Storage")
	storageClient.Invoke("add", "")
}
