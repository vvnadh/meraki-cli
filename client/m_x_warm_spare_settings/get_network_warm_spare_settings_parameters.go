// Code generated by go-swagger; DO NOT EDIT.

package m_x_warm_spare_settings

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

// NewGetNetworkWarmSpareSettingsParams creates a new GetNetworkWarmSpareSettingsParams object
// with the default values initialized.
func NewGetNetworkWarmSpareSettingsParams() *GetNetworkWarmSpareSettingsParams {
	var ()
	return &GetNetworkWarmSpareSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWarmSpareSettingsParamsWithTimeout creates a new GetNetworkWarmSpareSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkWarmSpareSettingsParamsWithTimeout(timeout time.Duration) *GetNetworkWarmSpareSettingsParams {
	var ()
	return &GetNetworkWarmSpareSettingsParams{

		timeout: timeout,
	}
}

// NewGetNetworkWarmSpareSettingsParamsWithContext creates a new GetNetworkWarmSpareSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkWarmSpareSettingsParamsWithContext(ctx context.Context) *GetNetworkWarmSpareSettingsParams {
	var ()
	return &GetNetworkWarmSpareSettingsParams{

		Context: ctx,
	}
}

// NewGetNetworkWarmSpareSettingsParamsWithHTTPClient creates a new GetNetworkWarmSpareSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkWarmSpareSettingsParamsWithHTTPClient(client *http.Client) *GetNetworkWarmSpareSettingsParams {
	var ()
	return &GetNetworkWarmSpareSettingsParams{
		HTTPClient: client,
	}
}

/*GetNetworkWarmSpareSettingsParams contains all the parameters to send to the API endpoint
for the get network warm spare settings operation typically these are written to a http.Request
*/
type GetNetworkWarmSpareSettingsParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network warm spare settings params
func (o *GetNetworkWarmSpareSettingsParams) WithTimeout(timeout time.Duration) *GetNetworkWarmSpareSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network warm spare settings params
func (o *GetNetworkWarmSpareSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network warm spare settings params
func (o *GetNetworkWarmSpareSettingsParams) WithContext(ctx context.Context) *GetNetworkWarmSpareSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network warm spare settings params
func (o *GetNetworkWarmSpareSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network warm spare settings params
func (o *GetNetworkWarmSpareSettingsParams) WithHTTPClient(client *http.Client) *GetNetworkWarmSpareSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network warm spare settings params
func (o *GetNetworkWarmSpareSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network warm spare settings params
func (o *GetNetworkWarmSpareSettingsParams) WithNetworkID(networkID string) *GetNetworkWarmSpareSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network warm spare settings params
func (o *GetNetworkWarmSpareSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWarmSpareSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
