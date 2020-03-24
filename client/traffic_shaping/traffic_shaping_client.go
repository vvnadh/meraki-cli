// Code generated by go-swagger; DO NOT EDIT.

package traffic_shaping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new traffic shaping API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for traffic shaping API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetNetworkSsidTrafficShaping(params *GetNetworkSsidTrafficShapingParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkSsidTrafficShapingOK, error)

	GetNetworkTrafficShaping(params *GetNetworkTrafficShapingParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkTrafficShapingOK, error)

	GetNetworkTrafficShapingApplicationCategories(params *GetNetworkTrafficShapingApplicationCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkTrafficShapingApplicationCategoriesOK, error)

	GetNetworkTrafficShapingDscpTaggingOptions(params *GetNetworkTrafficShapingDscpTaggingOptionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkTrafficShapingDscpTaggingOptionsOK, error)

	UpdateNetworkSsidTrafficShaping(params *UpdateNetworkSsidTrafficShapingParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkSsidTrafficShapingOK, error)

	UpdateNetworkTrafficShaping(params *UpdateNetworkTrafficShapingParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkTrafficShapingOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetNetworkSsidTrafficShaping gets network ssid traffic shaping

  Display the traffic shaping settings for a SSID on an MR network
*/
func (a *Client) GetNetworkSsidTrafficShaping(params *GetNetworkSsidTrafficShapingParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkSsidTrafficShapingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkSsidTrafficShapingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkSsidTrafficShaping",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/ssids/{number}/trafficShaping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkSsidTrafficShapingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkSsidTrafficShapingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkSsidTrafficShaping: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkTrafficShaping gets network traffic shaping

  Display the traffic shaping settings for an MX network
*/
func (a *Client) GetNetworkTrafficShaping(params *GetNetworkTrafficShapingParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkTrafficShapingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkTrafficShapingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkTrafficShaping",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/trafficShaping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkTrafficShapingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkTrafficShapingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkTrafficShaping: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkTrafficShapingApplicationCategories gets network traffic shaping application categories

  Returns the application categories for traffic shaping rules.
*/
func (a *Client) GetNetworkTrafficShapingApplicationCategories(params *GetNetworkTrafficShapingApplicationCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkTrafficShapingApplicationCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkTrafficShapingApplicationCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkTrafficShapingApplicationCategories",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/trafficShaping/applicationCategories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkTrafficShapingApplicationCategoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkTrafficShapingApplicationCategoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkTrafficShapingApplicationCategories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkTrafficShapingDscpTaggingOptions gets network traffic shaping dscp tagging options

  Returns the available DSCP tagging options for your traffic shaping rules.
*/
func (a *Client) GetNetworkTrafficShapingDscpTaggingOptions(params *GetNetworkTrafficShapingDscpTaggingOptionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkTrafficShapingDscpTaggingOptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkTrafficShapingDscpTaggingOptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkTrafficShapingDscpTaggingOptions",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/trafficShaping/dscpTaggingOptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkTrafficShapingDscpTaggingOptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkTrafficShapingDscpTaggingOptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkTrafficShapingDscpTaggingOptions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkSsidTrafficShaping updates network ssid traffic shaping

  Update the traffic shaping settings for an SSID on an MR network
*/
func (a *Client) UpdateNetworkSsidTrafficShaping(params *UpdateNetworkSsidTrafficShapingParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkSsidTrafficShapingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkSsidTrafficShapingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkSsidTrafficShaping",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/ssids/{number}/trafficShaping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkSsidTrafficShapingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkSsidTrafficShapingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkSsidTrafficShaping: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkTrafficShaping updates network traffic shaping

  Update the traffic shaping settings for an MX network
*/
func (a *Client) UpdateNetworkTrafficShaping(params *UpdateNetworkTrafficShapingParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkTrafficShapingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkTrafficShapingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkTrafficShaping",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/trafficShaping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkTrafficShapingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkTrafficShapingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkTrafficShaping: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
