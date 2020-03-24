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

// Type10 Type10
//
// The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'
//
// swagger:model Type10
type Type10 string

const (

	// Type10Text captures enum value "text"
	Type10Text Type10 = "text"

	// Type10IP captures enum value "ip"
	Type10IP Type10 = "ip"

	// Type10Hex captures enum value "hex"
	Type10Hex Type10 = "hex"

	// Type10Integer captures enum value "integer"
	Type10Integer Type10 = "integer"
)

// for schema
var type10Enum []interface{}

func init() {
	var res []Type10
	if err := json.Unmarshal([]byte(`["text","ip","hex","integer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		type10Enum = append(type10Enum, v)
	}
}

func (m Type10) validateType10Enum(path, location string, value Type10) error {
	if err := validate.Enum(path, location, value, type10Enum); err != nil {
		return err
	}
	return nil
}

// Validate validates this type10
func (m Type10) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateType10Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
