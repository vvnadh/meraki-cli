// Code generated by go-swagger; DO NOT EDIT.

package clients

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

// NewGetNetworkClientParams creates a new GetNetworkClientParams object
// with the default values initialized.
func NewGetNetworkClientParams() *GetNetworkClientParams {
	var ()
	return &GetNetworkClientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkClientParamsWithTimeout creates a new GetNetworkClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkClientParamsWithTimeout(timeout time.Duration) *GetNetworkClientParams {
	var ()
	return &GetNetworkClientParams{

		timeout: timeout,
	}
}

// NewGetNetworkClientParamsWithContext creates a new GetNetworkClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkClientParamsWithContext(ctx context.Context) *GetNetworkClientParams {
	var ()
	return &GetNetworkClientParams{

		Context: ctx,
	}
}

// NewGetNetworkClientParamsWithHTTPClient creates a new GetNetworkClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkClientParamsWithHTTPClient(client *http.Client) *GetNetworkClientParams {
	var ()
	return &GetNetworkClientParams{
		HTTPClient: client,
	}
}

/*GetNetworkClientParams contains all the parameters to send to the API endpoint
for the get network client operation typically these are written to a http.Request
*/
type GetNetworkClientParams struct {

	/*ClientID*/
	ClientID string
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network client params
func (o *GetNetworkClientParams) WithTimeout(timeout time.Duration) *GetNetworkClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network client params
func (o *GetNetworkClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network client params
func (o *GetNetworkClientParams) WithContext(ctx context.Context) *GetNetworkClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network client params
func (o *GetNetworkClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network client params
func (o *GetNetworkClientParams) WithHTTPClient(client *http.Client) *GetNetworkClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network client params
func (o *GetNetworkClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the get network client params
func (o *GetNetworkClientParams) WithClientID(clientID string) *GetNetworkClientParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get network client params
func (o *GetNetworkClientParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithNetworkID adds the networkID to the get network client params
func (o *GetNetworkClientParams) WithNetworkID(networkID string) *GetNetworkClientParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network client params
func (o *GetNetworkClientParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientId
	if err := r.SetPathParam("clientId", o.ClientID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
