// Code generated by go-swagger; DO NOT EDIT.

package switch_settings

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

// NewGetNetworkSwitchSettingsQosRulesOrderParams creates a new GetNetworkSwitchSettingsQosRulesOrderParams object
// with the default values initialized.
func NewGetNetworkSwitchSettingsQosRulesOrderParams() *GetNetworkSwitchSettingsQosRulesOrderParams {
	var ()
	return &GetNetworkSwitchSettingsQosRulesOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSwitchSettingsQosRulesOrderParamsWithTimeout creates a new GetNetworkSwitchSettingsQosRulesOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkSwitchSettingsQosRulesOrderParamsWithTimeout(timeout time.Duration) *GetNetworkSwitchSettingsQosRulesOrderParams {
	var ()
	return &GetNetworkSwitchSettingsQosRulesOrderParams{

		timeout: timeout,
	}
}

// NewGetNetworkSwitchSettingsQosRulesOrderParamsWithContext creates a new GetNetworkSwitchSettingsQosRulesOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkSwitchSettingsQosRulesOrderParamsWithContext(ctx context.Context) *GetNetworkSwitchSettingsQosRulesOrderParams {
	var ()
	return &GetNetworkSwitchSettingsQosRulesOrderParams{

		Context: ctx,
	}
}

// NewGetNetworkSwitchSettingsQosRulesOrderParamsWithHTTPClient creates a new GetNetworkSwitchSettingsQosRulesOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkSwitchSettingsQosRulesOrderParamsWithHTTPClient(client *http.Client) *GetNetworkSwitchSettingsQosRulesOrderParams {
	var ()
	return &GetNetworkSwitchSettingsQosRulesOrderParams{
		HTTPClient: client,
	}
}

/*GetNetworkSwitchSettingsQosRulesOrderParams contains all the parameters to send to the API endpoint
for the get network switch settings qos rules order operation typically these are written to a http.Request
*/
type GetNetworkSwitchSettingsQosRulesOrderParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network switch settings qos rules order params
func (o *GetNetworkSwitchSettingsQosRulesOrderParams) WithTimeout(timeout time.Duration) *GetNetworkSwitchSettingsQosRulesOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network switch settings qos rules order params
func (o *GetNetworkSwitchSettingsQosRulesOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network switch settings qos rules order params
func (o *GetNetworkSwitchSettingsQosRulesOrderParams) WithContext(ctx context.Context) *GetNetworkSwitchSettingsQosRulesOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network switch settings qos rules order params
func (o *GetNetworkSwitchSettingsQosRulesOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network switch settings qos rules order params
func (o *GetNetworkSwitchSettingsQosRulesOrderParams) WithHTTPClient(client *http.Client) *GetNetworkSwitchSettingsQosRulesOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network switch settings qos rules order params
func (o *GetNetworkSwitchSettingsQosRulesOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network switch settings qos rules order params
func (o *GetNetworkSwitchSettingsQosRulesOrderParams) WithNetworkID(networkID string) *GetNetworkSwitchSettingsQosRulesOrderParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network switch settings qos rules order params
func (o *GetNetworkSwitchSettingsQosRulesOrderParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSwitchSettingsQosRulesOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
