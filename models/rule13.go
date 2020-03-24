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

// Rule13 Rule13
//
// swagger:model Rule13
type Rule13 struct {

	// A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
	// Required: true
	Definitions []*Definition `json:"definitions"`

	// The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.
	//     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
	DscpTagValue int32 `json:"dscpTagValue,omitempty"`

	// per client bandwidth limits
	PerClientBandwidthLimits *PerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"`

	// A string, indicating the priority level for packets bound to your rule.
	//     Can be 'low', 'normal' or 'high'.
	Priority string `json:"priority,omitempty"`
}

// Validate validates this rule13
func (m *Rule13) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefinitions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerClientBandwidthLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rule13) validateDefinitions(formats strfmt.Registry) error {

	if err := validate.Required("definitions", "body", m.Definitions); err != nil {
		return err
	}

	for i := 0; i < len(m.Definitions); i++ {
		if swag.IsZero(m.Definitions[i]) { // not required
			continue
		}

		if m.Definitions[i] != nil {
			if err := m.Definitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("definitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Rule13) validatePerClientBandwidthLimits(formats strfmt.Registry) error {

	if swag.IsZero(m.PerClientBandwidthLimits) { // not required
		return nil
	}

	if m.PerClientBandwidthLimits != nil {
		if err := m.PerClientBandwidthLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perClientBandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rule13) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rule13) UnmarshalBinary(b []byte) error {
	var res Rule13
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
