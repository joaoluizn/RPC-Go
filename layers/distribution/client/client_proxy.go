package client

import "github.com/joaoluizn/RPC-go/network"

// NewClientProxy: Create ClientProxy instance;
func NewClientProxy(namingServerAddress string, serviceName string) *ClientProxy {
	return &ClientProxy{
		workerPool:   network.NewWorkerPool(namingServerAddress),
	}
}

// ClientProxy: Object reponsible for remote communication
type ClientProxy struct {
	workerPool   *network.WorkerPool
}

