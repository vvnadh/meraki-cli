// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApTagsAndVlanID ApTagsAndVlanId
//
// swagger:model ApTagsAndVlanId
type ApTagsAndVlanID struct {

	// Comma-separated list of AP tags
	Tags string `json:"tags,omitempty"`

	// Numerical identifier that is assigned to the VLAN
	VlanID int32 `json:"vlanId,omitempty"`
}

// Validate validates this ap tags and vlan Id
func (m *ApTagsAndVlanID) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApTagsAndVlanID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApTagsAndVlanID) UnmarshalBinary(b []byte) error {
	var res ApTagsAndVlanID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
