// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNetworkOneToOneNatRules updateNetworkOneToOneNatRules
//
// swagger:model updateNetworkOneToOneNatRules
type UpdateNetworkOneToOneNatRules struct {

	// An array of 1:1 nat rules
	// Required: true
	Rules []*Rule8 `json:"rules"`
}

// Validate validates this update network one to one nat rules
func (m *UpdateNetworkOneToOneNatRules) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkOneToOneNatRules) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("rules", "body", m.Rules); err != nil {
		return err
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
func (m *UpdateNetworkOneToOneNatRules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkOneToOneNatRules) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkOneToOneNatRules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
