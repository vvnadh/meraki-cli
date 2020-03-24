// Code generated by go-swagger; DO NOT EDIT.

package m_x_v_p_n_firewall

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateOrganizationVpnFirewallRulesReader is a Reader for the UpdateOrganizationVpnFirewallRules structure.
type UpdateOrganizationVpnFirewallRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationVpnFirewallRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationVpnFirewallRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOrganizationVpnFirewallRulesOK creates a UpdateOrganizationVpnFirewallRulesOK with default headers values
func NewUpdateOrganizationVpnFirewallRulesOK() *UpdateOrganizationVpnFirewallRulesOK {
	return &UpdateOrganizationVpnFirewallRulesOK{}
}

/*UpdateOrganizationVpnFirewallRulesOK handles this case with default header values.

Successful operation
*/
type UpdateOrganizationVpnFirewallRulesOK struct {
	Payload interface{}
}

func (o *UpdateOrganizationVpnFirewallRulesOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/vpnFirewallRules][%d] updateOrganizationVpnFirewallRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationVpnFirewallRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationVpnFirewallRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
