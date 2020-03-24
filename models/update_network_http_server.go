// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkHTTPServer updateNetworkHttpServer
//
// swagger:model updateNetworkHttpServer
type UpdateNetworkHTTPServer struct {

	// A name for easy reference to the HTTP server
	Name string `json:"name,omitempty"`

	// A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki.
	SharedSecret string `json:"sharedSecret,omitempty"`

	// The URL of the HTTP server
	URL string `json:"url,omitempty"`
}

// Validate validates this update network Http server
func (m *UpdateNetworkHTTPServer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkHTTPServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkHTTPServer) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkHTTPServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
