package infrastructure

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

// NewClientRequestHandler:  Create ClientRequestHandler instance;
func NewClientRequestHandler() *ClientRequestHandler {
	return &ClientRequestHandler{
		Client: &http.Client{
			// 10 seconds until timeout in any request.
			Timeout: 10 * time.Second},
	}
}

// ClientRequestHandler: Object used to send requests to a remote service;
type ClientRequestHandler struct {
	*http.Client
}

// Lookup looks for a remote service address for the naming service given
func (r *ClientRequestHandler) Lookup(namingServerAddr string, serviceName string) *http.Response {
	response, err := r.Get(
		// Lookup URL
		"http://" + namingServerAddr + "/lookup/" + serviceName,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	return response
}

// Send sends a invoke request to remote service
func (r *ClientRequestHandler) Send(remoteServiceAddr string, request *bytes.Buffer) *http.Response {
	response, err := r.Post(
		// URL
		"http://"+remoteServiceAddr+"/invoke/",
		// ContentType
		"service/json",
		// Data
		request,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	return response
}
