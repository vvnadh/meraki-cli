// Code generated by go-swagger; DO NOT EDIT.

package syslog_servers

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

// NewGetNetworkSyslogServersParams creates a new GetNetworkSyslogServersParams object
// with the default values initialized.
func NewGetNetworkSyslogServersParams() *GetNetworkSyslogServersParams {
	var ()
	return &GetNetworkSyslogServersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSyslogServersParamsWithTimeout creates a new GetNetworkSyslogServersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkSyslogServersParamsWithTimeout(timeout time.Duration) *GetNetworkSyslogServersParams {
	var ()
	return &GetNetworkSyslogServersParams{

		timeout: timeout,
	}
}

// NewGetNetworkSyslogServersParamsWithContext creates a new GetNetworkSyslogServersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkSyslogServersParamsWithContext(ctx context.Context) *GetNetworkSyslogServersParams {
	var ()
	return &GetNetworkSyslogServersParams{

		Context: ctx,
	}
}

// NewGetNetworkSyslogServersParamsWithHTTPClient creates a new GetNetworkSyslogServersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkSyslogServersParamsWithHTTPClient(client *http.Client) *GetNetworkSyslogServersParams {
	var ()
	return &GetNetworkSyslogServersParams{
		HTTPClient: client,
	}
}

/*GetNetworkSyslogServersParams contains all the parameters to send to the API endpoint
for the get network syslog servers operation typically these are written to a http.Request
*/
type GetNetworkSyslogServersParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network syslog servers params
func (o *GetNetworkSyslogServersParams) WithTimeout(timeout time.Duration) *GetNetworkSyslogServersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network syslog servers params
func (o *GetNetworkSyslogServersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network syslog servers params
func (o *GetNetworkSyslogServersParams) WithContext(ctx context.Context) *GetNetworkSyslogServersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network syslog servers params
func (o *GetNetworkSyslogServersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network syslog servers params
func (o *GetNetworkSyslogServersParams) WithHTTPClient(client *http.Client) *GetNetworkSyslogServersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network syslog servers params
func (o *GetNetworkSyslogServersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network syslog servers params
func (o *GetNetworkSyslogServersParams) WithNetworkID(networkID string) *GetNetworkSyslogServersParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network syslog servers params
func (o *GetNetworkSyslogServersParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSyslogServersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
