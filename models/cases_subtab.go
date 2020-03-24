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

// CasesSubtab CasesSubtab
//
// The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one
//     of 'default or inherit', 'hide' or 'show'.
//
// swagger:model CasesSubtab
type CasesSubtab string

const (

	// CasesSubtabDefaultOrInherit captures enum value "default or inherit"
	CasesSubtabDefaultOrInherit CasesSubtab = "default or inherit"

	// CasesSubtabHide captures enum value "hide"
	CasesSubtabHide CasesSubtab = "hide"

	// CasesSubtabShow captures enum value "show"
	CasesSubtabShow CasesSubtab = "show"
)

// for schema
var casesSubtabEnum []interface{}

func init() {
	var res []CasesSubtab
	if err := json.Unmarshal([]byte(`["default or inherit","hide","show"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		casesSubtabEnum = append(casesSubtabEnum, v)
	}
}

func (m CasesSubtab) validateCasesSubtabEnum(path, location string, value CasesSubtab) error {
	if err := validate.Enum(path, location, value, casesSubtabEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this cases subtab
func (m CasesSubtab) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCasesSubtabEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
