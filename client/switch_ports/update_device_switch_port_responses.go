// Code generated by go-swagger; DO NOT EDIT.

package switch_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateDeviceSwitchPortReader is a Reader for the UpdateDeviceSwitchPort structure.
type UpdateDeviceSwitchPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceSwitchPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceSwitchPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDeviceSwitchPortOK creates a UpdateDeviceSwitchPortOK with default headers values
func NewUpdateDeviceSwitchPortOK() *UpdateDeviceSwitchPortOK {
	return &UpdateDeviceSwitchPortOK{}
}

/*UpdateDeviceSwitchPortOK handles this case with default header values.

Successful operation
*/
type UpdateDeviceSwitchPortOK struct {
	Payload interface{}
}

func (o *UpdateDeviceSwitchPortOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/switchPorts/{number}][%d] updateDeviceSwitchPortOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceSwitchPortOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateDeviceSwitchPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
