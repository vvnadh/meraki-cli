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

// UpdateNetworkSmDeviceFieldsReader is a Reader for the UpdateNetworkSmDeviceFields structure.
type UpdateNetworkSmDeviceFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSmDeviceFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSmDeviceFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkSmDeviceFieldsOK creates a UpdateNetworkSmDeviceFieldsOK with default headers values
func NewUpdateNetworkSmDeviceFieldsOK() *UpdateNetworkSmDeviceFieldsOK {
	return &UpdateNetworkSmDeviceFieldsOK{}
}

/*UpdateNetworkSmDeviceFieldsOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkSmDeviceFieldsOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSmDeviceFieldsOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/sm/device/fields][%d] updateNetworkSmDeviceFieldsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSmDeviceFieldsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSmDeviceFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
