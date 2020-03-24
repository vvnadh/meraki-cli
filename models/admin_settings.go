// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminSettings AdminSettings
//
// Settings for describing which kinds of admins this policy applies to.
//
// swagger:model AdminSettings
type AdminSettings struct {

	// applies to
	AppliesTo AppliesTo `json:"appliesTo,omitempty"`

	// If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of
	//     entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of
	//     networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',
	//     specify an array of tag names.
	Values []string `json:"values"`
}

// Validate validates this admin settings
func (m *AdminSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppliesTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminSettings) validateAppliesTo(formats strfmt.Registry) error {

	if swag.IsZero(m.AppliesTo) { // not required
		return nil
	}

	if err := m.AppliesTo.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("appliesTo")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdminSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminSettings) UnmarshalBinary(b []byte) error {
	var res AdminSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
