// Code generated by go-swagger; DO NOT EDIT.

package content_filtering_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new content filtering rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for content filtering rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetNetworkContentFiltering(params *GetNetworkContentFilteringParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkContentFilteringOK, error)

	UpdateNetworkContentFiltering(params *UpdateNetworkContentFilteringParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkContentFilteringOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetNetworkContentFiltering gets network content filtering

  Return the content filtering settings for an MX network
*/
func (a *Client) GetNetworkContentFiltering(params *GetNetworkContentFilteringParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkContentFilteringOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkContentFilteringParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkContentFiltering",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/contentFiltering",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkContentFilteringReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkContentFilteringOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkContentFiltering: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkContentFiltering updates network content filtering

  Update the content filtering settings for an MX network
*/
func (a *Client) UpdateNetworkContentFiltering(params *UpdateNetworkContentFilteringParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkContentFilteringOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkContentFilteringParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkContentFiltering",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/contentFiltering",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkContentFilteringReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkContentFilteringOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkContentFiltering: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
