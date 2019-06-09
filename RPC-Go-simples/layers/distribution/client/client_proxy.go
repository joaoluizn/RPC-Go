package client

// NewClientProxy: Create ClientProxy instance;
func NewClientProxy(namingServerAddress string, serviceName string) *ClientProxy {
	return &ClientProxy{
		requestor: NewRequestor(namingServerAddress, serviceName),
	}
}

func (p *ClientProxy) UseRemoteService(op OperationArguments) {

	p.requestor.useRemoteService(op.OperationName, op.arg1, op.arg2)
}

// ClientProxy: Object reponsible for remote communication
type ClientProxy struct {
	requestor *Requestor
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
