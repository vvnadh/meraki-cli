// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkGroupPolicy updateNetworkGroupPolicy
//
// swagger:model updateNetworkGroupPolicy
type UpdateNetworkGroupPolicy struct {

	// bandwidth
	Bandwidth *Bandwidth `json:"bandwidth,omitempty"`

	// bonjour forwarding
	BonjourForwarding *BonjourForwarding `json:"bonjourForwarding,omitempty"`

	// content filtering
	ContentFiltering *ContentFiltering `json:"contentFiltering,omitempty"`

	// firewall and traffic shaping
	FirewallAndTrafficShaping *FirewallAndTrafficShaping `json:"firewallAndTrafficShaping,omitempty"`

	// The name for your group policy.
	Name string `json:"name,omitempty"`

	// scheduling
	Scheduling *Scheduling `json:"scheduling,omitempty"`

	// splash auth settings
	SplashAuthSettings SplashAuthSettings `json:"splashAuthSettings,omitempty"`

	// vlan tagging
	VlanTagging *VlanTagging `json:"vlanTagging,omitempty"`
}

// Validate validates this update network group policy
func (m *UpdateNetworkGroupPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBandwidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBonjourForwarding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentFiltering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallAndTrafficShaping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSplashAuthSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanTagging(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkGroupPolicy) validateBandwidth(formats strfmt.Registry) error {

	if swag.IsZero(m.Bandwidth) { // not required
		return nil
	}

	if m.Bandwidth != nil {
		if err := m.Bandwidth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bandwidth")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateNetworkGroupPolicy) validateBonjourForwarding(formats strfmt.Registry) error {

	if swag.IsZero(m.BonjourForwarding) { // not required
		return nil
	}

	if m.BonjourForwarding != nil {
		if err := m.BonjourForwarding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bonjourForwarding")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateNetworkGroupPolicy) validateContentFiltering(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentFiltering) { // not required
		return nil
	}

	if m.ContentFiltering != nil {
		if err := m.ContentFiltering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentFiltering")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateNetworkGroupPolicy) validateFirewallAndTrafficShaping(formats strfmt.Registry) error {

	if swag.IsZero(m.FirewallAndTrafficShaping) { // not required
		return nil
	}

	if m.FirewallAndTrafficShaping != nil {
		if err := m.FirewallAndTrafficShaping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firewallAndTrafficShaping")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateNetworkGroupPolicy) validateScheduling(formats strfmt.Registry) error {

	if swag.IsZero(m.Scheduling) { // not required
		return nil
	}

	if m.Scheduling != nil {
		if err := m.Scheduling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduling")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateNetworkGroupPolicy) validateSplashAuthSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.SplashAuthSettings) { // not required
		return nil
	}

	if err := m.SplashAuthSettings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("splashAuthSettings")
		}
		return err
	}

	return nil
}

func (m *UpdateNetworkGroupPolicy) validateVlanTagging(formats strfmt.Registry) error {

	if swag.IsZero(m.VlanTagging) { // not required
		return nil
	}

	if m.VlanTagging != nil {
		if err := m.VlanTagging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanTagging")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkGroupPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkGroupPolicy) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkGroupPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
