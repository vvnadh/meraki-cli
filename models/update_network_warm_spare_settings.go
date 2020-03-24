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

// UpdateNetworkWarmSpareSettings updateNetworkWarmSpareSettings
//
// swagger:model updateNetworkWarmSpareSettings
type UpdateNetworkWarmSpareSettings struct {

	// Enable warm spare
	// Required: true
	Enabled *bool `json:"enabled"`

	// Serial number of the warm spare appliance
	SpareSerial string `json:"spareSerial,omitempty"`

	// Uplink mode, either virtual or public
	UplinkMode string `json:"uplinkMode,omitempty"`

	// The WAN 1 shared IP
	VirtualIp1 string `json:"virtualIp1,omitempty"`

	// The WAN 2 shared IP
	VirtualIp2 string `json:"virtualIp2,omitempty"`
}

// Validate validates this update network warm spare settings
func (m *UpdateNetworkWarmSpareSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkWarmSpareSettings) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkWarmSpareSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkWarmSpareSettings) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWarmSpareSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
