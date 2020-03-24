// Code generated by go-swagger; DO NOT EDIT.

package net_flow_settings

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

// NewGetNetworkNetflowSettingsParams creates a new GetNetworkNetflowSettingsParams object
// with the default values initialized.
func NewGetNetworkNetflowSettingsParams() *GetNetworkNetflowSettingsParams {
	var ()
	return &GetNetworkNetflowSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkNetflowSettingsParamsWithTimeout creates a new GetNetworkNetflowSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkNetflowSettingsParamsWithTimeout(timeout time.Duration) *GetNetworkNetflowSettingsParams {
	var ()
	return &GetNetworkNetflowSettingsParams{

		timeout: timeout,
	}
}

// NewGetNetworkNetflowSettingsParamsWithContext creates a new GetNetworkNetflowSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkNetflowSettingsParamsWithContext(ctx context.Context) *GetNetworkNetflowSettingsParams {
	var ()
	return &GetNetworkNetflowSettingsParams{

		Context: ctx,
	}
}

// NewGetNetworkNetflowSettingsParamsWithHTTPClient creates a new GetNetworkNetflowSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkNetflowSettingsParamsWithHTTPClient(client *http.Client) *GetNetworkNetflowSettingsParams {
	var ()
	return &GetNetworkNetflowSettingsParams{
		HTTPClient: client,
	}
}

/*GetNetworkNetflowSettingsParams contains all the parameters to send to the API endpoint
for the get network netflow settings operation typically these are written to a http.Request
*/
type GetNetworkNetflowSettingsParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network netflow settings params
func (o *GetNetworkNetflowSettingsParams) WithTimeout(timeout time.Duration) *GetNetworkNetflowSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network netflow settings params
func (o *GetNetworkNetflowSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network netflow settings params
func (o *GetNetworkNetflowSettingsParams) WithContext(ctx context.Context) *GetNetworkNetflowSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network netflow settings params
func (o *GetNetworkNetflowSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network netflow settings params
func (o *GetNetworkNetflowSettingsParams) WithHTTPClient(client *http.Client) *GetNetworkNetflowSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network netflow settings params
func (o *GetNetworkNetflowSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network netflow settings params
func (o *GetNetworkNetflowSettingsParams) WithNetworkID(networkID string) *GetNetworkNetflowSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network netflow settings params
func (o *GetNetworkNetflowSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkNetflowSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
