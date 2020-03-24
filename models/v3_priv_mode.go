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

// V3PrivMode V3PrivMode
//
// The SNMP version 3 privacy mode. Can be either 'DES' or 'AES128'.
//
// swagger:model V3PrivMode
type V3PrivMode string

const (

	// V3PrivModeDES captures enum value "DES"
	V3PrivModeDES V3PrivMode = "DES"

	// V3PrivModeAES128 captures enum value "AES128"
	V3PrivModeAES128 V3PrivMode = "AES128"
)

// for schema
var v3PrivModeEnum []interface{}

func init() {
	var res []V3PrivMode
	if err := json.Unmarshal([]byte(`["DES","AES128"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3PrivModeEnum = append(v3PrivModeEnum, v)
	}
}

func (m V3PrivMode) validateV3PrivModeEnum(path, location string, value V3PrivMode) error {
	if err := validate.Enum(path, location, value, v3PrivModeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 priv mode
func (m V3PrivMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3PrivModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
