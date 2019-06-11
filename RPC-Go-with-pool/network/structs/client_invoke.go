package structs

// NewClientInvoke create ClientInvoke Entity
func NewClientInvoke(serviceName string, methodName string, arguments []interface{}) *ClientInvoke {
	return &ClientInvoke{ServiceName: serviceName, MethodName: methodName, Arguments: arguments}
}

// ClientInvoke Entity responsible to wrappers a client request to invoke a remote method
type ClientInvoke struct {
	ServiceName string
	MethodName  string
	Arguments   []interface{}
}
