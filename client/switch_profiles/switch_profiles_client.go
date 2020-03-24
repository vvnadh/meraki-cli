// Code generated by go-swagger; DO NOT EDIT.

package switch_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new switch profiles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for switch profiles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetOrganizationConfigTemplateSwitchProfiles(params *GetOrganizationConfigTemplateSwitchProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationConfigTemplateSwitchProfilesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetOrganizationConfigTemplateSwitchProfiles gets organization config template switch profiles

  List the switch profiles for your switch template configuration
*/
func (a *Client) GetOrganizationConfigTemplateSwitchProfiles(params *GetOrganizationConfigTemplateSwitchProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationConfigTemplateSwitchProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationConfigTemplateSwitchProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationConfigTemplateSwitchProfiles",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationId}/configTemplates/{configTemplateId}/switchProfiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationConfigTemplateSwitchProfilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationConfigTemplateSwitchProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationConfigTemplateSwitchProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}