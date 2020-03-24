// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkApplianceFirewallInboundFirewallRules updateNetworkApplianceFirewallInboundFirewallRules
//
// swagger:model updateNetworkApplianceFirewallInboundFirewallRules
type UpdateNetworkApplianceFirewallInboundFirewallRules struct {

	// An ordered array of the firewall rules (not including the default rule)
	Rules []*Rule1 `json:"rules"`

	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule bool `json:"syslogDefaultRule,omitempty"`
}

// Validate validates this update network appliance firewall inbound firewall rules
func (m *UpdateNetworkApplianceFirewallInboundFirewallRules) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkApplianceFirewallInboundFirewallRules) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkApplianceFirewallInboundFirewallRules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkApplianceFirewallInboundFirewallRules) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallInboundFirewallRules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
