// Code generated by go-swagger; DO NOT EDIT.

package m_x_inbound_firewall

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new m x inbound firewall API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for m x inbound firewall API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetNetworkApplianceFirewallInboundFirewallRules(params *GetNetworkApplianceFirewallInboundFirewallRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkApplianceFirewallInboundFirewallRulesOK, error)

	UpdateNetworkApplianceFirewallInboundFirewallRules(params *UpdateNetworkApplianceFirewallInboundFirewallRulesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkApplianceFirewallInboundFirewallRulesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetNetworkApplianceFirewallInboundFirewallRules gets network appliance firewall inbound firewall rules

  Return the inbound firewall rules for an MX network
*/
func (a *Client) GetNetworkApplianceFirewallInboundFirewallRules(params *GetNetworkApplianceFirewallInboundFirewallRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkApplianceFirewallInboundFirewallRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkApplianceFirewallInboundFirewallRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkApplianceFirewallInboundFirewallRules",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/appliance/firewall/inboundFirewallRules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkApplianceFirewallInboundFirewallRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkApplianceFirewallInboundFirewallRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkApplianceFirewallInboundFirewallRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkApplianceFirewallInboundFirewallRules updates network appliance firewall inbound firewall rules

  Update the inbound firewall rules of an MX network
*/
func (a *Client) UpdateNetworkApplianceFirewallInboundFirewallRules(params *UpdateNetworkApplianceFirewallInboundFirewallRulesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkApplianceFirewallInboundFirewallRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkApplianceFirewallInboundFirewallRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkApplianceFirewallInboundFirewallRules",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/appliance/firewall/inboundFirewallRules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkApplianceFirewallInboundFirewallRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkApplianceFirewallInboundFirewallRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkApplianceFirewallInboundFirewallRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
