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

// Protocol5 Protocol5
//
// TCP or UDP
//
// swagger:model Protocol5
type Protocol5 string

const (

	// Protocol5TCP captures enum value "tcp"
	Protocol5TCP Protocol5 = "tcp"

	// Protocol5UDP captures enum value "udp"
	Protocol5UDP Protocol5 = "udp"
)

// for schema
var protocol5Enum []interface{}

func init() {
	var res []Protocol5
	if err := json.Unmarshal([]byte(`["tcp","udp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protocol5Enum = append(protocol5Enum, v)
	}
}

func (m Protocol5) validateProtocol5Enum(path, location string, value Protocol5) error {
	if err := validate.Enum(path, location, value, protocol5Enum); err != nil {
		return err
	}
	return nil
}

// Validate validates this protocol5
func (m Protocol5) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProtocol5Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
