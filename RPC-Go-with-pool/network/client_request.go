package network

// NewClientRequest Create ClientRequest Entity;
func NewClientRequest(serviceName string, methodName string, args []interface{}) *ClientRequest {
	return &ClientRequest{ServiceName: serviceName, MethodName: methodName, Arguments: args}
}

// ClientRequest Entity responsible to wrap method information from client Requests;
type ClientRequest struct {
	ServiceName string
	MethodName  string
	Arguments   []interface{}
}
