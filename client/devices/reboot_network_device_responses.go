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

// RebootNetworkDeviceReader is a Reader for the RebootNetworkDevice structure.
type RebootNetworkDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RebootNetworkDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRebootNetworkDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRebootNetworkDeviceOK creates a RebootNetworkDeviceOK with default headers values
func NewRebootNetworkDeviceOK() *RebootNetworkDeviceOK {
	return &RebootNetworkDeviceOK{}
}

/*RebootNetworkDeviceOK handles this case with default header values.

Successful operation
*/
type RebootNetworkDeviceOK struct {
	Payload interface{}
}

func (o *RebootNetworkDeviceOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/devices/{serial}/reboot][%d] rebootNetworkDeviceOK  %+v", 200, o.Payload)
}

func (o *RebootNetworkDeviceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RebootNetworkDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
