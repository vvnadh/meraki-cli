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

// Mapping Mapping
//
// swagger:model Mapping
type Mapping struct {

	// The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.
	// Required: true
	Cos *int32 `json:"cos"`

	// The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.
	// Required: true
	Dscp *int32 `json:"dscp"`

	// Label for the mapping (optional).
	Title string `json:"title,omitempty"`
}

// Validate validates this mapping
func (m *Mapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDscp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mapping) validateCos(formats strfmt.Registry) error {

	if err := validate.Required("cos", "body", m.Cos); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validateDscp(formats strfmt.Registry) error {

	if err := validate.Required("dscp", "body", m.Dscp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Mapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Mapping) UnmarshalBinary(b []byte) error {
	var res Mapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
