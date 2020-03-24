// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PoliciesBySecurityAppliance PoliciesBySecurityAppliance
//
// An object, describing what the policy-connection association is for the security appliance. (Only relevant if the security appliance is actually within the network)
//
// swagger:model PoliciesBySecurityAppliance
type PoliciesBySecurityAppliance struct {

	// device policy
	DevicePolicy DevicePolicy1 `json:"devicePolicy,omitempty"`
}

// Validate validates this policies by security appliance
func (m *PoliciesBySecurityAppliance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevicePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoliciesBySecurityAppliance) validateDevicePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.DevicePolicy) { // not required
		return nil
	}

	if err := m.DevicePolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("devicePolicy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoliciesBySecurityAppliance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoliciesBySecurityAppliance) UnmarshalBinary(b []byte) error {
	var res PoliciesBySecurityAppliance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
