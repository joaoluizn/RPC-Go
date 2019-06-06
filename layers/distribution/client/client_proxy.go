package client

// NewClientProxy: Create ClientProxy instance;
func NewClientProxy(namingServerAddress string, serviceName string) *ClientProxy {
	return &ClientProxy{
		workerPool: NewWorkerPool(namingServerAddress, serviceName),
	}
}

// ClientProxy: Object reponsible for remote communication
type ClientProxy struct {
	workerPool *WorkerPool
}

func (p *ClientProxy) UseRemoteService(numOfOps int, clientOperations []string) {
	p.workerPool.useRemoteService(numOfOps, clientOperations)
}
