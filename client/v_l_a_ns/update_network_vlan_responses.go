// Code generated by go-swagger; DO NOT EDIT.

package v_l_a_ns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkVlanReader is a Reader for the UpdateNetworkVlan structure.
type UpdateNetworkVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkVlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkVlanOK creates a UpdateNetworkVlanOK with default headers values
func NewUpdateNetworkVlanOK() *UpdateNetworkVlanOK {
	return &UpdateNetworkVlanOK{}
}

/*UpdateNetworkVlanOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkVlanOK struct {
	Payload interface{}
}

func (o *UpdateNetworkVlanOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/vlans/{vlanId}][%d] updateNetworkVlanOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkVlanOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkVlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
