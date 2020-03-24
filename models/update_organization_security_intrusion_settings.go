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

// UpdateOrganizationSecurityIntrusionSettings updateOrganizationSecurityIntrusionSettings
//
// swagger:model updateOrganizationSecurityIntrusionSettings
type UpdateOrganizationSecurityIntrusionSettings struct {

	// Sets a list of specific SNORT® signatures to whitelist
	// Required: true
	WhitelistedRules []*WhitelistedRule `json:"whitelistedRules"`
}

// Validate validates this update organization security intrusion settings
func (m *UpdateOrganizationSecurityIntrusionSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWhitelistedRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateOrganizationSecurityIntrusionSettings) validateWhitelistedRules(formats strfmt.Registry) error {

	if err := validate.Required("whitelistedRules", "body", m.WhitelistedRules); err != nil {
		return err
	}

	for i := 0; i < len(m.WhitelistedRules); i++ {
		if swag.IsZero(m.WhitelistedRules[i]) { // not required
			continue
		}

		if m.WhitelistedRules[i] != nil {
			if err := m.WhitelistedRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("whitelistedRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateOrganizationSecurityIntrusionSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateOrganizationSecurityIntrusionSettings) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationSecurityIntrusionSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
