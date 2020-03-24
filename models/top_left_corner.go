// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TopLeftCorner TopLeftCorner
//
// The longitude and latitude of the top left corner of your floor plan.
//
// swagger:model TopLeftCorner
type TopLeftCorner struct {

	// Latitude
	Lat float64 `json:"lat,omitempty"`

	// Longitude
	Lng float64 `json:"lng,omitempty"`
}

// Validate validates this top left corner
func (m *TopLeftCorner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TopLeftCorner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopLeftCorner) UnmarshalBinary(b []byte) error {
	var res TopLeftCorner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
