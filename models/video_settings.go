// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VideoSettings VideoSettings
//
// Video quality and resolution settings for all the camera models.
//
// swagger:model VideoSettings
type VideoSettings struct {

	// m v12 m v22 m v72
	MV12MV22MV72 *MV12MV22MV72 `json:"MV12/MV22/MV72,omitempty"`

	// m v12 w e
	MV12WE *MV12WE `json:"MV12WE,omitempty"`

	// m v21 m v71
	MV21MV71 *MV21MV71 `json:"MV21/MV71,omitempty"`

	// m v22 x m v72 x
	MV22XMV72X *MV22XMV72X `json:"MV22X/MV72X,omitempty"`

	// m v32
	MV32 *MV32 `json:"MV32,omitempty"`
}

// Validate validates this video settings
func (m *VideoSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMV12MV22MV72(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMV12WE(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMV21MV71(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMV22XMV72X(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMV32(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VideoSettings) validateMV12MV22MV72(formats strfmt.Registry) error {

	if swag.IsZero(m.MV12MV22MV72) { // not required
		return nil
	}

	if m.MV12MV22MV72 != nil {
		if err := m.MV12MV22MV72.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MV12/MV22/MV72")
			}
			return err
		}
	}

	return nil
}

func (m *VideoSettings) validateMV12WE(formats strfmt.Registry) error {

	if swag.IsZero(m.MV12WE) { // not required
		return nil
	}

	if m.MV12WE != nil {
		if err := m.MV12WE.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MV12WE")
			}
			return err
		}
	}

	return nil
}

func (m *VideoSettings) validateMV21MV71(formats strfmt.Registry) error {

	if swag.IsZero(m.MV21MV71) { // not required
		return nil
	}

	if m.MV21MV71 != nil {
		if err := m.MV21MV71.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MV21/MV71")
			}
			return err
		}
	}

	return nil
}

func (m *VideoSettings) validateMV22XMV72X(formats strfmt.Registry) error {

	if swag.IsZero(m.MV22XMV72X) { // not required
		return nil
	}

	if m.MV22XMV72X != nil {
		if err := m.MV22XMV72X.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MV22X/MV72X")
			}
			return err
		}
	}

	return nil
}

func (m *VideoSettings) validateMV32(formats strfmt.Registry) error {

	if swag.IsZero(m.MV32) { // not required
		return nil
	}

	if m.MV32 != nil {
		if err := m.MV32.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MV32")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VideoSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VideoSettings) UnmarshalBinary(b []byte) error {
	var res VideoSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
