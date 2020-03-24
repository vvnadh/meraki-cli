// Code generated by go-swagger; DO NOT EDIT.

package firewalled_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkFirewalledServiceReader is a Reader for the UpdateNetworkFirewalledService structure.
type UpdateNetworkFirewalledServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkFirewalledServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkFirewalledServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkFirewalledServiceOK creates a UpdateNetworkFirewalledServiceOK with default headers values
func NewUpdateNetworkFirewalledServiceOK() *UpdateNetworkFirewalledServiceOK {
	return &UpdateNetworkFirewalledServiceOK{}
}

/*UpdateNetworkFirewalledServiceOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkFirewalledServiceOK struct {
	Payload interface{}
}

func (o *UpdateNetworkFirewalledServiceOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/firewalledServices/{service}][%d] updateNetworkFirewalledServiceOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkFirewalledServiceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkFirewalledServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
