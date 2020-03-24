// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkSwitchSettingsMtu updateNetworkSwitchSettingsMtu
//
// swagger:model updateNetworkSwitchSettingsMtu
type UpdateNetworkSwitchSettingsMtu struct {

	// MTU size for the entire network. Default value is 9578.
	DefaultMtuSize int32 `json:"defaultMtuSize,omitempty"`

	// Override MTU size for individual switches or switch profiles. An empty array will clear overrides.
	Overrides []*Override `json:"overrides"`
}

// Validate validates this update network switch settings mtu
func (m *UpdateNetworkSwitchSettingsMtu) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkSwitchSettingsMtu) validateOverrides(formats strfmt.Registry) error {

	if swag.IsZero(m.Overrides) { // not required
		return nil
	}

	for i := 0; i < len(m.Overrides); i++ {
		if swag.IsZero(m.Overrides[i]) { // not required
			continue
		}

		if m.Overrides[i] != nil {
			if err := m.Overrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkSwitchSettingsMtu) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkSwitchSettingsMtu) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchSettingsMtu
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}