package client

import (
	"log"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/network"
)

// NewWorkerPool create a new worker pool to handle services requests
func NewWorkerPool(namingServerAddress string, serviceName string) *WorkerPool {
	return &WorkerPool{
		requestor:   NewRequestor(namingServerAddress),
		serviceName: serviceName,
	}
}

// WorkerPool Object reponsible for handling workers;
type WorkerPool struct {
	requestor   *Requestor
	serviceName string
}

// Operation Object responsible to wrap service call data
type Operation struct {
	operationName string
	args1         interface{}
	args2         interface{}
	operationId   int
}

// Response Object responsible to unwrap response for data compliance puposes
type Response struct {
	operationResponse interface{}
	operationId       int
}

// Invoke Run desired method on remote server;
func (w *WorkerPool) Invoke(operation Operation) network.Response {
	return w.requestor.Invoke(w.serviceName, operation.operationName, operation.args1, operation.args2)
}

// useRemoteService Responsible for remote service calling by worker pool;
func (w *WorkerPool) useRemoteService(numOfOps int, opName []string, opArgs1 []interface{}, opArgs2 []interface{}) {

	clientOperations := make([]Operation, numOfOps)
	for j := 0; j < numOfOps; j++ {
		clientOperations[j] = Operation{opName[j], opArgs1[j], opArgs2[j], j}
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
	counter += numOfOps
	for workCounter < numOfOps {

	}
}

var (
	counter     = 0
	workCounter = 0
)

// UseRemoteService Call a Remote Procedure
func (w *WorkerPool) UseRemoteService(operations <-chan Operation, responses chan<- Response) {
	for op := range operations {
		log.Printf("Operation %d: Calling Remote Procedure: '%s'", op.operationId+counter, op.operationName)
		// This Invoke can receive the operation to be executed and arguments needed
		responses <- Response{(w.Invoke(op).Content[0]), op.operationId + counter}
		workCounter++
	}
}
