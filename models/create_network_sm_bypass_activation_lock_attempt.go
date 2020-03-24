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

// CreateNetworkSmBypassActivationLockAttempt createNetworkSmBypassActivationLockAttempt
//
// swagger:model createNetworkSmBypassActivationLockAttempt
type CreateNetworkSmBypassActivationLockAttempt struct {

	// The ids of the devices to attempt activation lock bypass.
	// Required: true
	Ids []string `json:"ids"`
}

// Validate validates this create network sm bypass activation lock attempt
func (m *CreateNetworkSmBypassActivationLockAttempt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateNetworkSmBypassActivationLockAttempt) validateIds(formats strfmt.Registry) error {

	if err := validate.Required("ids", "body", m.Ids); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateNetworkSmBypassActivationLockAttempt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateNetworkSmBypassActivationLockAttempt) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSmBypassActivationLockAttempt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
