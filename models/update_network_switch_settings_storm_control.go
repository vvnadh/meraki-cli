// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkSwitchSettingsStormControl updateNetworkSwitchSettingsStormControl
//
// swagger:model updateNetworkSwitchSettingsStormControl
type UpdateNetworkSwitchSettingsStormControl struct {

	// Percentage (1 to 99) of total available port bandwidth for broadcast traffic type. Default value 100 percent rate is to clear the configuration.
	BroadcastThreshold int32 `json:"broadcastThreshold,omitempty"`

	// Percentage (1 to 99) of total available port bandwidth for multicast traffic type. Default value 100 percent rate is to clear the configuration.
	MulticastThreshold int32 `json:"multicastThreshold,omitempty"`

	// Percentage (1 to 99) of total available port bandwidth for unknown unicast (dlf-destination lookup failure) traffic type. Default value 100 percent rate is to clear the configuration.
	UnknownUnicastThreshold int32 `json:"unknownUnicastThreshold,omitempty"`
}

// Validate validates this update network switch settings storm control
func (m *UpdateNetworkSwitchSettingsStormControl) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkSwitchSettingsStormControl) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkSwitchSettingsStormControl) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchSettingsStormControl
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
