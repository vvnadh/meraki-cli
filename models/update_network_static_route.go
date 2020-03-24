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

// UpdateNetworkStaticRoute updateNetworkStaticRoute
//
// swagger:model updateNetworkStaticRoute
type UpdateNetworkStaticRoute struct {

	// The enabled state of the static route
	Enabled bool `json:"enabled,omitempty"`

	// The DHCP fixed IP assignments on the static route. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.
	FixedIPAssignments interface{} `json:"fixedIpAssignments,omitempty"`

	// The gateway IP (next hop) of the static route
	GatewayIP string `json:"gatewayIp,omitempty"`

	// The name of the static route
	Name string `json:"name,omitempty"`

	// The DHCP reserved IP ranges on the static route
	ReservedIPRanges []*ReservedIPRange1 `json:"reservedIpRanges"`

	// The subnet of the static route
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this update network static route
func (m *UpdateNetworkStaticRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReservedIPRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkStaticRoute) validateReservedIPRanges(formats strfmt.Registry) error {

	if swag.IsZero(m.ReservedIPRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.ReservedIPRanges); i++ {
		if swag.IsZero(m.ReservedIPRanges[i]) { // not required
			continue
		}

		if m.ReservedIPRanges[i] != nil {
			if err := m.ReservedIPRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reservedIpRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkStaticRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkStaticRoute) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkStaticRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
