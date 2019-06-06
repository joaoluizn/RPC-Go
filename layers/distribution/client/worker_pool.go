package client

import (
	"log"
	"time"

	"github.com/joaoluizn/RPC-Go/network"
)

func NewWorkerPool(namingServerAddress string, serviceName string) *WorkerPool {
	return &WorkerPool{
		requestor:   NewRequestor(namingServerAddress),
		serviceName: serviceName,
	}
}

// ClientProxy: Object reponsible for remote communication
type WorkerPool struct {
	requestor   *Requestor
	serviceName string
}

type Operation struct {
	operationName string
	operationId   int
}

type Response struct {
	operationResponse interface{}
	operationId       int
}

// Invoke: Run desired method on remote server;
func (w *WorkerPool) Invoke(methodName string, arguments ...interface{}) network.Response {
	return w.requestor.Invoke(w.serviceName, methodName, arguments)
}

func (w *WorkerPool) useRemoteService(numOfOps int, clientOps []string) {

	clientOperations := make([]Operation, numOfOps)
	for j := 0; j < numOfOps; j++ {
		clientOperations[j] = Operation{clientOps[j], j}
	}

	operations := make(chan Operation, numOfOps)
	responses := make(chan Response, numOfOps)

	go w.UseRemoteService(operations, responses)
	go w.UseRemoteService(operations, responses)
	go w.UseRemoteService(operations, responses)
	go w.UseRemoteService(operations, responses)
	go w.UseRemoteService(operations, responses)

	for i := 0; i < numOfOps; i++ {
		operations <- clientOperations[i]
	}
	close(operations)

	for r := 0; r < numOfOps; r++ {
		response := <-responses
		log.Printf("Response of Operation %d: %s", response.operationId, response.operationResponse)
	}
}

func (w *WorkerPool) UseRemoteService(operations <-chan Operation, responses chan<- Response) {

	time.Sleep(time.Second * 2)

	for op := range operations {
		// Calling a Remote Procedure
		log.Printf("Operation %d: Calling Remote Procedure: '%s'", op.operationId, op.operationName)
		// This Invoke can receive the operation to be executed and arguments needed
		responses <- Response{(w.Invoke(op.operationName, "teste", op.operationId).Content[0]), op.operationId}
	}
}
