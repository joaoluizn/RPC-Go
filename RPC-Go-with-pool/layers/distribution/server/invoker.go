package server

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/network"
	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/network/structs"
	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/services/storage"
)

// NewInvoker creates Invoker Entity;
func NewInvoker() *Invoker {
	return &Invoker{
		RemoteService: storage.NewRemoteService(),
		marshaller:    network.NewMarshaller(),
	}
}

// Invoker Entity responsible for run requested methods;
type Invoker struct {
	RemoteService *storage.RemoteService
	marshaller    *network.Marshaller
}

// Invoke Invoke a requested method;
func (i *Invoker) Invoke(request *http.Request) []byte {
	clientInvoke := i.marshaller.UnmarshalClientInvokeRequest(request)
	output := i.invoke(clientInvoke)
	return i.marshaller.MarshalClientResponse(output)
}

//invoke Invoke a requested method;
func (i *Invoker) invoke(clientInvoke *structs.ClientInvoke) interface{} {
	log.Printf("Invoking: %s.%s(%s: R$%s)",
		clientInvoke.ServiceName, clientInvoke.MethodName, clientInvoke.Arguments[0], fmt.Sprintf("%.2f", clientInvoke.Arguments[1]),
	)
	service := i.getService(clientInvoke.ServiceName)
	method := i.getMethod(service, clientInvoke.MethodName)
	arguments := i.getArguments(method, clientInvoke.Arguments)
	outputs := method.Call(arguments)
	return i.getMethodReturn(outputs)
}

// getService Get service from invocation request;
func (i *Invoker) getService(serviceName string) reflect.Value {
	serviceValue := reflect.ValueOf(i.RemoteService.GetService(serviceName))
	if !serviceValue.IsValid() {
		log.Fatalf("Could not find Service: '%s'", serviceName)
	}

	return serviceValue
}

// getMethod Get method from invocation request;
func (i *Invoker) getMethod(service reflect.Value, methodName string) reflect.Value {
	methodReflectionValue := service.MethodByName(methodName)

	if !methodReflectionValue.IsValid() {
		log.Fatalf("Could not find Method: '%s' in Service: '%s'", methodName, service.Type().String())
	}

	return methodReflectionValue
}

// getArguments Get Arguments from invocation request
func (i *Invoker) getArguments(method reflect.Value, args []interface{}) []reflect.Value {
	argsValue := make([]reflect.Value, len(args))

	for index := range argsValue {
		arg := args[index]
		var newArg interface{}

		switch method.Type().In(index).Kind() {
		case reflect.Int:
			newArg = int(arg.(float64))
		case reflect.Float64:
			newArg = arg.(float64)
		case reflect.String:
			newArg = arg.(string)
		case reflect.Slice:
			newArg = arg
		}

		argsValue[index] = reflect.ValueOf(newArg)
	}

	return argsValue
}

// getMethodReturn Get method return from invocation calls
func (i *Invoker) getMethodReturn(outputs []reflect.Value) interface{} {
	outputsInterface := make([]interface{}, len(outputs))
	for index := range outputsInterface {
		outputsInterface[index] = outputs[index].Interface()
	}

	return outputsInterface
}
