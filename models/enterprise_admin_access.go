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

// EnterpriseAdminAccess EnterpriseAdminAccess
//
// Whether or not an SSID is accessible by 'enterprise' administrators ('access disabled' or 'access enabled')
//
// swagger:model EnterpriseAdminAccess
type EnterpriseAdminAccess string

const (

	// EnterpriseAdminAccessAccessDisabled captures enum value "access disabled"
	EnterpriseAdminAccessAccessDisabled EnterpriseAdminAccess = "access disabled"

	// EnterpriseAdminAccessAccessEnabled captures enum value "access enabled"
	EnterpriseAdminAccessAccessEnabled EnterpriseAdminAccess = "access enabled"
)

// for schema
var enterpriseAdminAccessEnum []interface{}

func init() {
	var res []EnterpriseAdminAccess
	if err := json.Unmarshal([]byte(`["access disabled","access enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enterpriseAdminAccessEnum = append(enterpriseAdminAccessEnum, v)
	}
}

func (m EnterpriseAdminAccess) validateEnterpriseAdminAccessEnum(path, location string, value EnterpriseAdminAccess) error {
	if err := validate.Enum(path, location, value, enterpriseAdminAccessEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this enterprise admin access
func (m EnterpriseAdminAccess) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEnterpriseAdminAccessEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
