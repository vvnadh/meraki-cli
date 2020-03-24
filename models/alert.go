// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Alert Alert
//
// swagger:model Alert
type Alert struct {

	// A hash of destinations for this specific alert. Keys include: emails: A list of emails that will recieve information about the alert, allAdmins: If true, then all network admins will receive emails, and snmp: If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.
	AlertDestinations interface{} `json:"alertDestinations,omitempty"`

	// A boolean depicting if the alert is turned on or off
	Enabled bool `json:"enabled,omitempty"`

	// A hash of specific configuration data for the alert. Only filters specific to the alert will be updated.
	Filters interface{} `json:"filters,omitempty"`

	// The type of alert
	Type string `json:"type,omitempty"`
}

// Validate validates this alert
func (m *Alert) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Alert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alert) UnmarshalBinary(b []byte) error {
	var res Alert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
