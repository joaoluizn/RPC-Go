package client

func main() {
	storageClient := NewClientProxy("0.0.0.0:8923", "Storage")
	storageClient.Invoke("add", "")
}
