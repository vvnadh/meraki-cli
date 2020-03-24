// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Rule1 Rule1
//
// swagger:model Rule1
type Rule1 struct {

	// Description of the rule (optional)
	Comment string `json:"comment,omitempty"`

	// Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	// Required: true
	DestCidr *string `json:"destCidr"`

	// Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	DestPort string `json:"destPort,omitempty"`

	// policy
	// Required: true
	Policy Policy `json:"policy"`

	// protocol
	// Required: true
	Protocol Protocol `json:"protocol"`

	// Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	// Required: true
	SrcCidr *string `json:"srcCidr"`

	// Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SrcPort string `json:"srcPort,omitempty"`

	// Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
	SyslogEnabled bool `json:"syslogEnabled,omitempty"`
}

// Validate validates this rule1
func (m *Rule1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rule1) validateDestCidr(formats strfmt.Registry) error {

	if err := validate.Required("destCidr", "body", m.DestCidr); err != nil {
		return err
	}

	return nil
}

func (m *Rule1) validatePolicy(formats strfmt.Registry) error {

	if err := m.Policy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policy")
		}
		return err
	}

	return nil
}

func (m *Rule1) validateProtocol(formats strfmt.Registry) error {

	if err := m.Protocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protocol")
		}
		return err
	}

	return nil
}

func (m *Rule1) validateSrcCidr(formats strfmt.Registry) error {

	if err := validate.Required("srcCidr", "body", m.SrcCidr); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rule1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rule1) UnmarshalBinary(b []byte) error {
	var res Rule1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
