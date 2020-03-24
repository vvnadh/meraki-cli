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

// UpdateNetworkSmDevicesTags updateNetworkSmDevicesTags
//
// swagger:model updateNetworkSmDevicesTags
type UpdateNetworkSmDevicesTags struct {

	// The ids of the devices to be modified.
	Ids string `json:"ids,omitempty"`

	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be modified.
	Scope string `json:"scope,omitempty"`

	// The serials of the devices to be modified.
	Serials string `json:"serials,omitempty"`

	// The tags to be added, deleted, or updated.
	// Required: true
	Tags *string `json:"tags"`

	// One of add, delete, or update. Only devices that have been modified will be returned.
	// Required: true
	UpdateAction *string `json:"updateAction"`

	// The wifiMacs of the devices to be modified.
	WifiMacs string `json:"wifiMacs,omitempty"`
}

// Validate validates this update network sm devices tags
func (m *UpdateNetworkSmDevicesTags) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkSmDevicesTags) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNetworkSmDevicesTags) validateUpdateAction(formats strfmt.Registry) error {

	if err := validate.Required("updateAction", "body", m.UpdateAction); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkSmDevicesTags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkSmDevicesTags) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSmDevicesTags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
