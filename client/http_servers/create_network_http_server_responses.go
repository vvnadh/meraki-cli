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

// CreateNetworkHTTPServerReader is a Reader for the CreateNetworkHTTPServer structure.
type CreateNetworkHTTPServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkHTTPServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkHTTPServerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateNetworkHTTPServerCreated creates a CreateNetworkHTTPServerCreated with default headers values
func NewCreateNetworkHTTPServerCreated() *CreateNetworkHTTPServerCreated {
	return &CreateNetworkHTTPServerCreated{}
}

/*CreateNetworkHTTPServerCreated handles this case with default header values.

Successful operation
*/
type CreateNetworkHTTPServerCreated struct {
	Payload interface{}
}

func (o *CreateNetworkHTTPServerCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/httpServers][%d] createNetworkHttpServerCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkHTTPServerCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkHTTPServerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
