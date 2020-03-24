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

// MoveOrganizationLicensesSeats moveOrganizationLicensesSeats
//
// swagger:model moveOrganizationLicensesSeats
type MoveOrganizationLicensesSeats struct {

	// The ID of the organization to move the SM seats to
	// Required: true
	DestOrganizationID *string `json:"destOrganizationId"`

	// The ID of the SM license to move the seats from
	// Required: true
	LicenseID *string `json:"licenseId"`

	// The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license
	// Required: true
	SeatCount *int32 `json:"seatCount"`
}

// Validate validates this move organization licenses seats
func (m *MoveOrganizationLicensesSeats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeatCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveOrganizationLicensesSeats) validateDestOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("destOrganizationId", "body", m.DestOrganizationID); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrganizationLicensesSeats) validateLicenseID(formats strfmt.Registry) error {

	if err := validate.Required("licenseId", "body", m.LicenseID); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrganizationLicensesSeats) validateSeatCount(formats strfmt.Registry) error {

	if err := validate.Required("seatCount", "body", m.SeatCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveOrganizationLicensesSeats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveOrganizationLicensesSeats) UnmarshalBinary(b []byte) error {
	var res MoveOrganizationLicensesSeats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
