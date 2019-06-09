package network

// NewClientRequest: Create ClientRequest instance;
func NewClientRequest(serviceName string, methodName string, args []interface{}) *ClientRequest {
	return &ClientRequest{ServiceName: serviceName, MethodName: methodName, Arguments: args}
}

// ClientRequest: Object that stores client request to execute a remote method;
type ClientRequest struct {
	ServiceName string
	MethodName  string
	Arguments   []interface{}
}
