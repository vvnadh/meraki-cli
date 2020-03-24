// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CycleDeviceSwitchPortsReader is a Reader for the CycleDeviceSwitchPorts structure.
type CycleDeviceSwitchPortsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CycleDeviceSwitchPortsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCycleDeviceSwitchPortsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCycleDeviceSwitchPortsOK creates a CycleDeviceSwitchPortsOK with default headers values
func NewCycleDeviceSwitchPortsOK() *CycleDeviceSwitchPortsOK {
	return &CycleDeviceSwitchPortsOK{}
}

/*CycleDeviceSwitchPortsOK handles this case with default header values.

Successful operation
*/
type CycleDeviceSwitchPortsOK struct {
	Payload interface{}
}

func (o *CycleDeviceSwitchPortsOK) Error() string {
	return fmt.Sprintf("[POST /devices/{serial}/switch/ports/cycle][%d] cycleDeviceSwitchPortsOK  %+v", 200, o.Payload)
}

func (o *CycleDeviceSwitchPortsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CycleDeviceSwitchPortsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
