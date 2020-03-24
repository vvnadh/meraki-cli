// Code generated by go-swagger; DO NOT EDIT.

package p_i_i

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateNetworkPiiRequestReader is a Reader for the CreateNetworkPiiRequest structure.
type CreateNetworkPiiRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkPiiRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkPiiRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateNetworkPiiRequestCreated creates a CreateNetworkPiiRequestCreated with default headers values
func NewCreateNetworkPiiRequestCreated() *CreateNetworkPiiRequestCreated {
	return &CreateNetworkPiiRequestCreated{}
}

/*CreateNetworkPiiRequestCreated handles this case with default header values.

Successful operation
*/
type CreateNetworkPiiRequestCreated struct {
	Payload interface{}
}

func (o *CreateNetworkPiiRequestCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/pii/requests][%d] createNetworkPiiRequestCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkPiiRequestCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkPiiRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
