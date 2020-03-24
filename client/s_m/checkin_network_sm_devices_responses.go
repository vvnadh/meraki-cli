// Code generated by go-swagger; DO NOT EDIT.

package s_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CheckinNetworkSmDevicesReader is a Reader for the CheckinNetworkSmDevices structure.
type CheckinNetworkSmDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckinNetworkSmDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckinNetworkSmDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckinNetworkSmDevicesOK creates a CheckinNetworkSmDevicesOK with default headers values
func NewCheckinNetworkSmDevicesOK() *CheckinNetworkSmDevicesOK {
	return &CheckinNetworkSmDevicesOK{}
}

/*CheckinNetworkSmDevicesOK handles this case with default header values.

Successful operation
*/
type CheckinNetworkSmDevicesOK struct {
	Payload interface{}
}

func (o *CheckinNetworkSmDevicesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/sm/devices/checkin][%d] checkinNetworkSmDevicesOK  %+v", 200, o.Payload)
}

func (o *CheckinNetworkSmDevicesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckinNetworkSmDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
