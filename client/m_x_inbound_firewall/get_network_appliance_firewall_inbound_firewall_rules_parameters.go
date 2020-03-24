// Code generated by go-swagger; DO NOT EDIT.

package m_x_inbound_firewall

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetNetworkApplianceFirewallInboundFirewallRulesParams creates a new GetNetworkApplianceFirewallInboundFirewallRulesParams object
// with the default values initialized.
func NewGetNetworkApplianceFirewallInboundFirewallRulesParams() *GetNetworkApplianceFirewallInboundFirewallRulesParams {
	var ()
	return &GetNetworkApplianceFirewallInboundFirewallRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkApplianceFirewallInboundFirewallRulesParamsWithTimeout creates a new GetNetworkApplianceFirewallInboundFirewallRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkApplianceFirewallInboundFirewallRulesParamsWithTimeout(timeout time.Duration) *GetNetworkApplianceFirewallInboundFirewallRulesParams {
	var ()
	return &GetNetworkApplianceFirewallInboundFirewallRulesParams{

		timeout: timeout,
	}
}

// NewGetNetworkApplianceFirewallInboundFirewallRulesParamsWithContext creates a new GetNetworkApplianceFirewallInboundFirewallRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkApplianceFirewallInboundFirewallRulesParamsWithContext(ctx context.Context) *GetNetworkApplianceFirewallInboundFirewallRulesParams {
	var ()
	return &GetNetworkApplianceFirewallInboundFirewallRulesParams{

		Context: ctx,
	}
}

// NewGetNetworkApplianceFirewallInboundFirewallRulesParamsWithHTTPClient creates a new GetNetworkApplianceFirewallInboundFirewallRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkApplianceFirewallInboundFirewallRulesParamsWithHTTPClient(client *http.Client) *GetNetworkApplianceFirewallInboundFirewallRulesParams {
	var ()
	return &GetNetworkApplianceFirewallInboundFirewallRulesParams{
		HTTPClient: client,
	}
}

/*GetNetworkApplianceFirewallInboundFirewallRulesParams contains all the parameters to send to the API endpoint
for the get network appliance firewall inbound firewall rules operation typically these are written to a http.Request
*/
type GetNetworkApplianceFirewallInboundFirewallRulesParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network appliance firewall inbound firewall rules params
func (o *GetNetworkApplianceFirewallInboundFirewallRulesParams) WithTimeout(timeout time.Duration) *GetNetworkApplianceFirewallInboundFirewallRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network appliance firewall inbound firewall rules params
func (o *GetNetworkApplianceFirewallInboundFirewallRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network appliance firewall inbound firewall rules params
func (o *GetNetworkApplianceFirewallInboundFirewallRulesParams) WithContext(ctx context.Context) *GetNetworkApplianceFirewallInboundFirewallRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network appliance firewall inbound firewall rules params
func (o *GetNetworkApplianceFirewallInboundFirewallRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network appliance firewall inbound firewall rules params
func (o *GetNetworkApplianceFirewallInboundFirewallRulesParams) WithHTTPClient(client *http.Client) *GetNetworkApplianceFirewallInboundFirewallRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network appliance firewall inbound firewall rules params
func (o *GetNetworkApplianceFirewallInboundFirewallRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network appliance firewall inbound firewall rules params
func (o *GetNetworkApplianceFirewallInboundFirewallRulesParams) WithNetworkID(networkID string) *GetNetworkApplianceFirewallInboundFirewallRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network appliance firewall inbound firewall rules params
func (o *GetNetworkApplianceFirewallInboundFirewallRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkApplianceFirewallInboundFirewallRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
