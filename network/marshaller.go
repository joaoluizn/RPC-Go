package network

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joaoluizn/RPC-go/network/structs"
)

// NewMarshaller build new instance of marshaller
func NewMarshaller() *Marshaller {
	return &Marshaller{}
}

// Marshaller handles frequest / response data serialization and deserialization
type Marshaller struct {
}

// Naming Service Marshaller
// UnmarshallLookupResponse: Deserializes Naming Service Server Response;
func (m *Marshaller) MarshallLookupResponse(address string) []byte {
	addressBytes, err := json.Marshal(address)
	if err != nil {
		log.Fatal(err.Error())
	}

	return addressBytes
}

func (m *Marshaller) UnmarshallLookupResponse(httpResponse *http.Response) string {
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var content string
	json.Unmarshal(body, &content)

	return content
}

func (m *Marshaller) MarshallRegistrationResponse(messages []string) []byte {
	messagesBytes, err := json.Marshal(messages)
	if err != nil {
		log.Fatal(err.Error())
	}
	return messagesBytes
}

func (m *Marshaller) UnMarshallRegistrationResponse(httpResponse *http.Response) []string {
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var content []string
	json.Unmarshal(body, &content)

	return content
}

func (m *Marshaller) UnmarshalClientInvokeRequest(htttpRequest *http.Request) *structs.ClientInvoke {
	body, err := ioutil.ReadAll(htttpRequest.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var invokeRequest *structs.ClientInvoke
	json.Unmarshal(body, &invokeRequest)

	return invokeRequest
}

// MarshallClientRequest: Serializes a client request;
func (m *Marshaller) MarshallClientRequest(clientInvokeRequest *ClientRequest) *bytes.Buffer {
	requestBytes, err := json.Marshal(clientInvokeRequest)
	if err != nil {
		log.Fatal(err.Error())
	}

	return bytes.NewBuffer(requestBytes)
}

// UnmarshallClientResponse: Deserializes a response to the client;
func (m *Marshaller) UnmarshallClientResponse(httpResponse *http.Response) Response {
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var contentResponse Response
	json.Unmarshal(body, &contentResponse.Content)

	return contentResponse
}

// Those marshall functions will handle new operation registration over the naming service
// MarshalNamingServiceRegistration: serializes namingServiceRegistration Object
func (m *Marshaller) MarshalNamingServiceRegistration(namingServiceRegistration *NamingServiceRegistration) *bytes.Buffer {
	objectBytes, err := json.Marshal(namingServiceRegistration)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("MarshalNamingServiceRegistration: Marshall complete: Service Registration data ready to be sent")
	return bytes.NewBuffer(objectBytes)
}

// UnmarshalNamingServiceRegistration deserializes a request
func (m *Marshaller) UnmarshalNamingServiceRegistration(httpRequest *http.Request) *NamingServiceRegistration {
	var registrationRequest NamingServiceRegistration
	body, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(body, &registrationRequest)
	return &registrationRequest
}

// MarshalClientResponse serializes a response
func (m *Marshaller) MarshalClientResponse(response interface{}) []byte {
	responseByte, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	return responseByte
}
