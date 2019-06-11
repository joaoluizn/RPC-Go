package network

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joaoluizn/RPC-Go/RPC-Go-with-pool/network/structs"
)

// NewMarshaller Create marshaller Entity
func NewMarshaller() *Marshaller {
	return &Marshaller{}
}

// Marshaller handles request / response data serialization and deserialization
type Marshaller struct{}

// Naming Service Marshall

// MarshallLookupResponse Serialize Lookup Response;
func (m *Marshaller) MarshallLookupResponse(address string) []byte {
	addressBytes, err := json.Marshal(address)
	if err != nil {
		log.Fatal(err.Error())
	}

	return addressBytes
}

// UnmarshallLookupResponse Deserialize Lookup Response;
func (m *Marshaller) UnmarshallLookupResponse(httpResponse *http.Response) string {
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var content string
	json.Unmarshal(body, &content)

	return content
}

// MarshallRegistrationResponse Serialize Registration Response;
func (m *Marshaller) MarshallRegistrationResponse(messages []string) []byte {
	messagesBytes, err := json.Marshal(messages)
	if err != nil {
		log.Fatal(err.Error())
	}
	return messagesBytes
}

// UnMarshallRegistrationResponse Deserialize Registration Response;
func (m *Marshaller) UnMarshallRegistrationResponse(httpResponse *http.Response) []string {
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var content []string
	json.Unmarshal(body, &content)

	return content
}

// Invoke Marshall

// UnmarshalClientInvokeRequest Deserialize Client Invoke Request;
func (m *Marshaller) UnmarshalClientInvokeRequest(htttpRequest *http.Request) *structs.ClientInvoke {
	body, err := ioutil.ReadAll(htttpRequest.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var invokeRequest *structs.ClientInvoke
	json.Unmarshal(body, &invokeRequest)

	return invokeRequest
}

// MarshallClientRequest Serialize Client Request;
func (m *Marshaller) MarshallClientRequest(clientInvokeRequest *ClientRequest) *bytes.Buffer {
	requestBytes, err := json.Marshal(clientInvokeRequest)
	if err != nil {
		log.Fatal(err.Error())
	}

	return bytes.NewBuffer(requestBytes)
}

// UnmarshallClientResponse Deserializes a response to the client;
func (m *Marshaller) UnmarshallClientResponse(httpResponse *http.Response) Response {
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var contentResponse Response
	json.Unmarshal(body, &contentResponse.Content)

	return contentResponse
}

// Naming Service Marshall

// MarshalNamingServiceRegistration Serializes namingServiceRegistration Object
func (m *Marshaller) MarshalNamingServiceRegistration(namingServiceRegistration *NamingServiceRegistration) *bytes.Buffer {
	objectBytes, err := json.Marshal(namingServiceRegistration)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("MarshalNamingServiceRegistration: Marshall complete: Service Registration data ready to be sent")
	return bytes.NewBuffer(objectBytes)
}

// UnmarshalNamingServiceRegistration Deserializes request to Service Registration
func (m *Marshaller) UnmarshalNamingServiceRegistration(httpRequest *http.Request) *NamingServiceRegistration {
	var registrationRequest NamingServiceRegistration
	body, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(body, &registrationRequest)
	return &registrationRequest
}

// MarshalClientResponse serializes Client response
func (m *Marshaller) MarshalClientResponse(response interface{}) []byte {
	responseByte, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	return responseByte
}
