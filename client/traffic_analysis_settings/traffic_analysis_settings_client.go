// Code generated by go-swagger; DO NOT EDIT.

package traffic_analysis_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new traffic analysis settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for traffic analysis settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetNetworkTrafficAnalysisSettings(params *GetNetworkTrafficAnalysisSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkTrafficAnalysisSettingsOK, error)

	UpdateNetworkTrafficAnalysisSettings(params *UpdateNetworkTrafficAnalysisSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkTrafficAnalysisSettingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetNetworkTrafficAnalysisSettings gets network traffic analysis settings

  Return the traffic analysis settings for a network
*/
func (a *Client) GetNetworkTrafficAnalysisSettings(params *GetNetworkTrafficAnalysisSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkTrafficAnalysisSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkTrafficAnalysisSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkTrafficAnalysisSettings",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/trafficAnalysisSettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkTrafficAnalysisSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkTrafficAnalysisSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkTrafficAnalysisSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkTrafficAnalysisSettings updates network traffic analysis settings

  Update the traffic analysis settings for a network
*/
func (a *Client) UpdateNetworkTrafficAnalysisSettings(params *UpdateNetworkTrafficAnalysisSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkTrafficAnalysisSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkTrafficAnalysisSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkTrafficAnalysisSettings",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/trafficAnalysisSettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkTrafficAnalysisSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkTrafficAnalysisSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkTrafficAnalysisSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
