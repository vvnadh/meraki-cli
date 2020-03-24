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

// Rule9 Rule9
//
// swagger:model Rule9
type Rule9 struct {

	// An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any)
	// Required: true
	AllowedIps []string `json:"allowedIps"`

	// The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	// Required: true
	LanIP *string `json:"lanIp"`

	// A port or port ranges that will receive the forwarded traffic from the WAN
	// Required: true
	LocalPort *string `json:"localPort"`

	// A descriptive name for the rule
	Name string `json:"name,omitempty"`

	// protocol
	// Required: true
	Protocol Protocol5 `json:"protocol"`

	// A port or port ranges that will be forwarded to the host on the LAN
	// Required: true
	PublicPort *string `json:"publicPort"`

	// uplink
	Uplink Uplink3 `json:"uplink,omitempty"`
}

// Validate validates this rule9
func (m *Rule9) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUplink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rule9) validateAllowedIps(formats strfmt.Registry) error {

	if err := validate.Required("allowedIps", "body", m.AllowedIps); err != nil {
		return err
	}

	return nil
}

func (m *Rule9) validateLanIP(formats strfmt.Registry) error {

	if err := validate.Required("lanIp", "body", m.LanIP); err != nil {
		return err
	}

	return nil
}

func (m *Rule9) validateLocalPort(formats strfmt.Registry) error {

	if err := validate.Required("localPort", "body", m.LocalPort); err != nil {
		return err
	}

	return nil
}

func (m *Rule9) validateProtocol(formats strfmt.Registry) error {

	if err := m.Protocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protocol")
		}
		return err
	}

	return nil
}

func (m *Rule9) validatePublicPort(formats strfmt.Registry) error {

	if err := validate.Required("publicPort", "body", m.PublicPort); err != nil {
		return err
	}

	return nil
}

func (m *Rule9) validateUplink(formats strfmt.Registry) error {

	if swag.IsZero(m.Uplink) { // not required
		return nil
	}

	if err := m.Uplink.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("uplink")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rule9) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rule9) UnmarshalBinary(b []byte) error {
	var res Rule9
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
