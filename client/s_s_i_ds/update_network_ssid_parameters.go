// Code generated by go-swagger; DO NOT EDIT.

package s_s_i_ds

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

	"github.com/cisco-sso/meraki-cli/models"
)

// NewUpdateNetworkSsidParams creates a new UpdateNetworkSsidParams object
// with the default values initialized.
func NewUpdateNetworkSsidParams() *UpdateNetworkSsidParams {
	var ()
	return &UpdateNetworkSsidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSsidParamsWithTimeout creates a new UpdateNetworkSsidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkSsidParamsWithTimeout(timeout time.Duration) *UpdateNetworkSsidParams {
	var ()
	return &UpdateNetworkSsidParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkSsidParamsWithContext creates a new UpdateNetworkSsidParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkSsidParamsWithContext(ctx context.Context) *UpdateNetworkSsidParams {
	var ()
	return &UpdateNetworkSsidParams{

		Context: ctx,
	}
}

// NewUpdateNetworkSsidParamsWithHTTPClient creates a new UpdateNetworkSsidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkSsidParamsWithHTTPClient(client *http.Client) *UpdateNetworkSsidParams {
	var ()
	return &UpdateNetworkSsidParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkSsidParams contains all the parameters to send to the API endpoint
for the update network ssid operation typically these are written to a http.Request
*/
type UpdateNetworkSsidParams struct {

	/*NetworkID*/
	NetworkID string
	/*Number*/
	Number string
	/*UpdateNetworkSsid*/
	UpdateNetworkSsid *models.UpdateNetworkSsid

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network ssid params
func (o *UpdateNetworkSsidParams) WithTimeout(timeout time.Duration) *UpdateNetworkSsidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network ssid params
func (o *UpdateNetworkSsidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network ssid params
func (o *UpdateNetworkSsidParams) WithContext(ctx context.Context) *UpdateNetworkSsidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network ssid params
func (o *UpdateNetworkSsidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network ssid params
func (o *UpdateNetworkSsidParams) WithHTTPClient(client *http.Client) *UpdateNetworkSsidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network ssid params
func (o *UpdateNetworkSsidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network ssid params
func (o *UpdateNetworkSsidParams) WithNetworkID(networkID string) *UpdateNetworkSsidParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network ssid params
func (o *UpdateNetworkSsidParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the update network ssid params
func (o *UpdateNetworkSsidParams) WithNumber(number string) *UpdateNetworkSsidParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the update network ssid params
func (o *UpdateNetworkSsidParams) SetNumber(number string) {
	o.Number = number
}

// WithUpdateNetworkSsid adds the updateNetworkSsid to the update network ssid params
func (o *UpdateNetworkSsidParams) WithUpdateNetworkSsid(updateNetworkSsid *models.UpdateNetworkSsid) *UpdateNetworkSsidParams {
	o.SetUpdateNetworkSsid(updateNetworkSsid)
	return o
}

// SetUpdateNetworkSsid adds the updateNetworkSsid to the update network ssid params
func (o *UpdateNetworkSsidParams) SetUpdateNetworkSsid(updateNetworkSsid *models.UpdateNetworkSsid) {
	o.UpdateNetworkSsid = updateNetworkSsid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSsidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param number
	if err := r.SetPathParam("number", o.Number); err != nil {
		return err
	}

	if o.UpdateNetworkSsid != nil {
		if err := r.SetBodyParam(o.UpdateNetworkSsid); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}