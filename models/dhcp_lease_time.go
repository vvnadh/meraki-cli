// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DhcpLeaseTime DhcpLeaseTime
//
// The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'
//
// swagger:model DhcpLeaseTime
type DhcpLeaseTime string

const (

	// DhcpLeaseTimeNr30Minutes captures enum value "30 minutes"
	DhcpLeaseTimeNr30Minutes DhcpLeaseTime = "30 minutes"

	// DhcpLeaseTimeNr1Hour captures enum value "1 hour"
	DhcpLeaseTimeNr1Hour DhcpLeaseTime = "1 hour"

	// DhcpLeaseTimeNr4Hours captures enum value "4 hours"
	DhcpLeaseTimeNr4Hours DhcpLeaseTime = "4 hours"

	// DhcpLeaseTimeNr12Hours captures enum value "12 hours"
	DhcpLeaseTimeNr12Hours DhcpLeaseTime = "12 hours"

	// DhcpLeaseTimeNr1Day captures enum value "1 day"
	DhcpLeaseTimeNr1Day DhcpLeaseTime = "1 day"

	// DhcpLeaseTimeNr1Week captures enum value "1 week"
	DhcpLeaseTimeNr1Week DhcpLeaseTime = "1 week"
)

// for schema
var dhcpLeaseTimeEnum []interface{}

func init() {
	var res []DhcpLeaseTime
	if err := json.Unmarshal([]byte(`["30 minutes","1 hour","4 hours","12 hours","1 day","1 week"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dhcpLeaseTimeEnum = append(dhcpLeaseTimeEnum, v)
	}
}

func (m DhcpLeaseTime) validateDhcpLeaseTimeEnum(path, location string, value DhcpLeaseTime) error {
	if err := validate.Enum(path, location, value, dhcpLeaseTimeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this dhcp lease time
func (m DhcpLeaseTime) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDhcpLeaseTimeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
