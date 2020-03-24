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

// ReservedIPRange1 ReservedIpRange1
//
// swagger:model ReservedIpRange1
type ReservedIPRange1 struct {

	// A text comment for the reserved range
	// Required: true
	Comment *string `json:"comment"`

	// The last IP in the reserved range
	// Required: true
	End *string `json:"end"`

	// The first IP in the reserved range
	// Required: true
	Start *string `json:"start"`
}

// Validate validates this reserved Ip range1
func (m *ReservedIPRange1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReservedIPRange1) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *ReservedIPRange1) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	return nil
}

func (m *ReservedIPRange1) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReservedIPRange1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReservedIPRange1) UnmarshalBinary(b []byte) error {
	var res ReservedIPRange1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
