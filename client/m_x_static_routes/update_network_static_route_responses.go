// Code generated by go-swagger; DO NOT EDIT.

package m_x_static_routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkStaticRouteReader is a Reader for the UpdateNetworkStaticRoute structure.
type UpdateNetworkStaticRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkStaticRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkStaticRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkStaticRouteOK creates a UpdateNetworkStaticRouteOK with default headers values
func NewUpdateNetworkStaticRouteOK() *UpdateNetworkStaticRouteOK {
	return &UpdateNetworkStaticRouteOK{}
}

/*UpdateNetworkStaticRouteOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkStaticRouteOK struct {
	Payload interface{}
}

func (o *UpdateNetworkStaticRouteOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/staticRoutes/{staticRouteId}][%d] updateNetworkStaticRouteOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkStaticRouteOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkStaticRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
