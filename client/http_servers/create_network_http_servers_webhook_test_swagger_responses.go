// Code generated by go-swagger; DO NOT EDIT.

package http_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateNetworkHTTPServersWebhookTestReader is a Reader for the CreateNetworkHTTPServersWebhookTest structure.
type CreateNetworkHTTPServersWebhookTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkHTTPServersWebhookTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNetworkHTTPServersWebhookTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateNetworkHTTPServersWebhookTestOK creates a CreateNetworkHTTPServersWebhookTestOK with default headers values
func NewCreateNetworkHTTPServersWebhookTestOK() *CreateNetworkHTTPServersWebhookTestOK {
	return &CreateNetworkHTTPServersWebhookTestOK{}
}

/*CreateNetworkHTTPServersWebhookTestOK handles this case with default header values.

Successful operation
*/
type CreateNetworkHTTPServersWebhookTestOK struct {
	Payload interface{}
}

func (o *CreateNetworkHTTPServersWebhookTestOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/httpServers/webhookTests][%d] createNetworkHttpServersWebhookTestOK  %+v", 200, o.Payload)
}

func (o *CreateNetworkHTTPServersWebhookTestOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkHTTPServersWebhookTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
