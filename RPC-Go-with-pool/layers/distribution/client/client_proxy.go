package client

// NewClientProxy: Create ClientProxy instance;
func NewClientProxy(namingServerAddress string, serviceName string) *ClientProxy {
	return &ClientProxy{
		workerPool: NewWorkerPool(namingServerAddress, serviceName),
	}
}

func (p *ClientProxy) UseRemoteService(numOfOps int, clientOperations []OperationArguments) {
	operationNames := make([]string, numOfOps)
	operationArgs1 := make([]interface{}, numOfOps)
	operationArgs2 := make([]interface{}, numOfOps)
	for i := 0; i < numOfOps; i++ {
		operationNames[i] = clientOperations[i].OperationName
		operationArgs1[i] = clientOperations[i].arg1
		operationArgs2[i] = clientOperations[i].arg2
	}
	p.workerPool.useRemoteService(numOfOps, operationNames, operationArgs1, operationArgs2)
}

// ClientProxy: Object reponsible for remote communication
type ClientProxy struct {
	workerPool *WorkerPool
}

type OperationArguments struct {
	OperationName string
	arg1          interface{}
	arg2          interface{}
}

func NewOperation(OperationName string, arg1 interface{}, arg2 interface{}) OperationArguments {
	return OperationArguments{
		OperationName: OperationName,
		arg1:          arg1,
		arg2:          arg2,
	}
}
