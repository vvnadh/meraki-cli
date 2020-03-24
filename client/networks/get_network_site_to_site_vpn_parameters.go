// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewGetNetworkSiteToSiteVpnParams creates a new GetNetworkSiteToSiteVpnParams object
// with the default values initialized.
func NewGetNetworkSiteToSiteVpnParams() *GetNetworkSiteToSiteVpnParams {
	var ()
	return &GetNetworkSiteToSiteVpnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSiteToSiteVpnParamsWithTimeout creates a new GetNetworkSiteToSiteVpnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkSiteToSiteVpnParamsWithTimeout(timeout time.Duration) *GetNetworkSiteToSiteVpnParams {
	var ()
	return &GetNetworkSiteToSiteVpnParams{

		timeout: timeout,
	}
}

// NewGetNetworkSiteToSiteVpnParamsWithContext creates a new GetNetworkSiteToSiteVpnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkSiteToSiteVpnParamsWithContext(ctx context.Context) *GetNetworkSiteToSiteVpnParams {
	var ()
	return &GetNetworkSiteToSiteVpnParams{

		Context: ctx,
	}
}

// NewGetNetworkSiteToSiteVpnParamsWithHTTPClient creates a new GetNetworkSiteToSiteVpnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkSiteToSiteVpnParamsWithHTTPClient(client *http.Client) *GetNetworkSiteToSiteVpnParams {
	var ()
	return &GetNetworkSiteToSiteVpnParams{
		HTTPClient: client,
	}
}

/*GetNetworkSiteToSiteVpnParams contains all the parameters to send to the API endpoint
for the get network site to site vpn operation typically these are written to a http.Request
*/
type GetNetworkSiteToSiteVpnParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network site to site vpn params
func (o *GetNetworkSiteToSiteVpnParams) WithTimeout(timeout time.Duration) *GetNetworkSiteToSiteVpnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network site to site vpn params
func (o *GetNetworkSiteToSiteVpnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network site to site vpn params
func (o *GetNetworkSiteToSiteVpnParams) WithContext(ctx context.Context) *GetNetworkSiteToSiteVpnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network site to site vpn params
func (o *GetNetworkSiteToSiteVpnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network site to site vpn params
func (o *GetNetworkSiteToSiteVpnParams) WithHTTPClient(client *http.Client) *GetNetworkSiteToSiteVpnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network site to site vpn params
func (o *GetNetworkSiteToSiteVpnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network site to site vpn params
func (o *GetNetworkSiteToSiteVpnParams) WithNetworkID(networkID string) *GetNetworkSiteToSiteVpnParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network site to site vpn params
func (o *GetNetworkSiteToSiteVpnParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSiteToSiteVpnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
