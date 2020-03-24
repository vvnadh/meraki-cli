// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateOrganizationThirdPartyVPNPeers updateOrganizationThirdPartyVPNPeers
//
// swagger:model updateOrganizationThirdPartyVPNPeers
type UpdateOrganizationThirdPartyVPNPeers struct {

	// The list of VPN peers
	// Required: true
	Peers []*Peer `json:"peers"`
}

// Validate validates this update organization third party v p n peers
func (m *UpdateOrganizationThirdPartyVPNPeers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateOrganizationThirdPartyVPNPeers) validatePeers(formats strfmt.Registry) error {

	if err := validate.Required("peers", "body", m.Peers); err != nil {
		return err
	}

	for i := 0; i < len(m.Peers); i++ {
		if swag.IsZero(m.Peers[i]) { // not required
			continue
		}

		if m.Peers[i] != nil {
			if err := m.Peers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateOrganizationThirdPartyVPNPeers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateOrganizationThirdPartyVPNPeers) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationThirdPartyVPNPeers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
