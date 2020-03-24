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

// RenewOrganizationLicensesSeats renewOrganizationLicensesSeats
//
// swagger:model renewOrganizationLicensesSeats
type RenewOrganizationLicensesSeats struct {

	// The ID of the SM license to renew. This license must already be assigned to an SM network
	// Required: true
	LicenseIDToRenew *string `json:"licenseIdToRenew"`

	// The SM license to use to renew the seats on 'licenseIdToRenew'. This license must have at least as many seats available as there are seats on 'licenseIdToRenew'
	// Required: true
	UnusedLicenseID *string `json:"unusedLicenseId"`
}

// Validate validates this renew organization licenses seats
func (m *RenewOrganizationLicensesSeats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLicenseIDToRenew(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnusedLicenseID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RenewOrganizationLicensesSeats) validateLicenseIDToRenew(formats strfmt.Registry) error {

	if err := validate.Required("licenseIdToRenew", "body", m.LicenseIDToRenew); err != nil {
		return err
	}

	return nil
}

func (m *RenewOrganizationLicensesSeats) validateUnusedLicenseID(formats strfmt.Registry) error {

	if err := validate.Required("unusedLicenseId", "body", m.UnusedLicenseID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RenewOrganizationLicensesSeats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RenewOrganizationLicensesSeats) UnmarshalBinary(b []byte) error {
	var res RenewOrganizationLicensesSeats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
