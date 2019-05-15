package main

import (
	// "log"
	"../distribution"
)

func main(){
	storageClient := distribution.NewClientProxy("0.0.0.0:8550", "Storage")
	storageClient.Invoke("add", "")
}