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

// UpdateNetworkSmDeviceFields updateNetworkSmDeviceFields
//
// swagger:model updateNetworkSmDeviceFields
type UpdateNetworkSmDeviceFields struct {

	// device fields
	// Required: true
	DeviceFields *DeviceFields `json:"deviceFields"`

	// The id of the device to be modified.
	ID string `json:"id,omitempty"`

	// The serial of the device to be modified.
	Serial string `json:"serial,omitempty"`

	// The wifiMac of the device to be modified.
	WifiMac string `json:"wifiMac,omitempty"`
}

// Validate validates this update network sm device fields
func (m *UpdateNetworkSmDeviceFields) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkSmDeviceFields) validateDeviceFields(formats strfmt.Registry) error {

	if err := validate.Required("deviceFields", "body", m.DeviceFields); err != nil {
		return err
	}

	if m.DeviceFields != nil {
		if err := m.DeviceFields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceFields")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkSmDeviceFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkSmDeviceFields) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSmDeviceFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
