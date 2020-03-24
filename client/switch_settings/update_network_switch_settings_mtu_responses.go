// Code generated by go-swagger; DO NOT EDIT.

package switch_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkSwitchSettingsMtuReader is a Reader for the UpdateNetworkSwitchSettingsMtu structure.
type UpdateNetworkSwitchSettingsMtuReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchSettingsMtuReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchSettingsMtuOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkSwitchSettingsMtuOK creates a UpdateNetworkSwitchSettingsMtuOK with default headers values
func NewUpdateNetworkSwitchSettingsMtuOK() *UpdateNetworkSwitchSettingsMtuOK {
	return &UpdateNetworkSwitchSettingsMtuOK{}
}

/*UpdateNetworkSwitchSettingsMtuOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkSwitchSettingsMtuOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSwitchSettingsMtuOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/settings/mtu][%d] updateNetworkSwitchSettingsMtuOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchSettingsMtuOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchSettingsMtuOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
