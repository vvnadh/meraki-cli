// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WhitelistedRule WhitelistedRule
//
// swagger:model WhitelistedRule
type WhitelistedRule struct {

	// Message is optional and is ignored on a PUT call. It is allowed in order for PUT to be compatible with GET
	Message string `json:"message,omitempty"`

	// A rule identifier of the format meraki:intrusion/snort/GID/<gid>/SID/<sid>. gid and sid can be obtained from either https://www.snort.org/rule-docs or as ruleIds from the security events in /organization/[orgId]/securityEvents
	// Required: true
	RuleID *string `json:"ruleId"`
}

// Validate validates this whitelisted rule
func (m *WhitelistedRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WhitelistedRule) validateRuleID(formats strfmt.Registry) error {

	if err := validate.Required("ruleId", "body", m.RuleID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WhitelistedRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WhitelistedRule) UnmarshalBinary(b []byte) error {
	var res WhitelistedRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
