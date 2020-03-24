// Code generated by go-swagger; DO NOT EDIT.

package m_g_port_forwarding_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateDeviceCellularGatewaySettingsPortForwardingRulesReader is a Reader for the UpdateDeviceCellularGatewaySettingsPortForwardingRules structure.
type UpdateDeviceCellularGatewaySettingsPortForwardingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesOK creates a UpdateDeviceCellularGatewaySettingsPortForwardingRulesOK with default headers values
func NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesOK() *UpdateDeviceCellularGatewaySettingsPortForwardingRulesOK {
	return &UpdateDeviceCellularGatewaySettingsPortForwardingRulesOK{}
}

/*UpdateDeviceCellularGatewaySettingsPortForwardingRulesOK handles this case with default header values.

Successful operation
*/
type UpdateDeviceCellularGatewaySettingsPortForwardingRulesOK struct {
	Payload interface{}
}

func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/cellularGateway/settings/portForwardingRules][%d] updateDeviceCellularGatewaySettingsPortForwardingRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
