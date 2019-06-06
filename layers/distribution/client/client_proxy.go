package client

// NewClientProxy: Create ClientProxy instance;
func NewClientProxy(namingServerAddress string, serviceName string) *ClientProxy {
	return &ClientProxy{
		workerPool: NewWorkerPool(namingServerAddress, serviceName),
	}
}

func (p *ClientProxy) UseRemoteService(numOfOps int, clientOperations []OperationArguments) {
	operationNames := make([]string, numOfOps)
	operationArg1 := make([]interface{}, numOfOps)
	operationArg2 := make([]interface{}, numOfOps)
	p.workerPool.useRemoteService(numOfOps, operationNames, operationArg1, operationArg2)
}

// ClientProxy: Object reponsible for remote communication
type ClientProxy struct {
	workerPool *WorkerPool
}

type OperationArguments struct {
	OperationName string
	args1         interface{}
	args2         interface{}
}

func (p *ClientProxy) NewOperation(OperationName string, args1 interface{}, args2 interface{}) *OperationArguments {
	return &OperationArguments{
		OperationName: OperationName,
		args1:         args1,
		args2:         args2,
	}
}
