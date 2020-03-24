// Code generated by go-swagger; DO NOT EDIT.

package group_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateNetworkGroupPolicy(params *CreateNetworkGroupPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNetworkGroupPolicyCreated, error)

	DeleteNetworkGroupPolicy(params *DeleteNetworkGroupPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworkGroupPolicyNoContent, error)

	GetNetworkGroupPolicies(params *GetNetworkGroupPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkGroupPoliciesOK, error)

	GetNetworkGroupPolicy(params *GetNetworkGroupPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkGroupPolicyOK, error)

	UpdateNetworkGroupPolicy(params *UpdateNetworkGroupPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkGroupPolicyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateNetworkGroupPolicy creates network group policy

  Create a group policy
*/
func (a *Client) CreateNetworkGroupPolicy(params *CreateNetworkGroupPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNetworkGroupPolicyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkGroupPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNetworkGroupPolicy",
		Method:             "POST",
		PathPattern:        "/networks/{networkId}/groupPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNetworkGroupPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateNetworkGroupPolicyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createNetworkGroupPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNetworkGroupPolicy deletes network group policy

  Delete a group policy
*/
func (a *Client) DeleteNetworkGroupPolicy(params *DeleteNetworkGroupPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworkGroupPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkGroupPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNetworkGroupPolicy",
		Method:             "DELETE",
		PathPattern:        "/networks/{networkId}/groupPolicies/{groupPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNetworkGroupPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetworkGroupPolicyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNetworkGroupPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkGroupPolicies gets network group policies

  List the group policies in a network
*/
func (a *Client) GetNetworkGroupPolicies(params *GetNetworkGroupPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkGroupPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkGroupPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkGroupPolicies",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/groupPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkGroupPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkGroupPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkGroupPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkGroupPolicy gets network group policy

  Display a group policy
*/
func (a *Client) GetNetworkGroupPolicy(params *GetNetworkGroupPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkGroupPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkGroupPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkGroupPolicy",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/groupPolicies/{groupPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkGroupPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkGroupPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkGroupPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkGroupPolicy updates network group policy

  Update a group policy
*/
func (a *Client) UpdateNetworkGroupPolicy(params *UpdateNetworkGroupPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkGroupPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkGroupPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkGroupPolicy",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/groupPolicies/{groupPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkGroupPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkGroupPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkGroupPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
