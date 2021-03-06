// Code generated by go-swagger; DO NOT EDIT.

package content_filtering_rules

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

// NewGetNetworkContentFilteringParams creates a new GetNetworkContentFilteringParams object
// with the default values initialized.
func NewGetNetworkContentFilteringParams() *GetNetworkContentFilteringParams {
	var ()
	return &GetNetworkContentFilteringParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkContentFilteringParamsWithTimeout creates a new GetNetworkContentFilteringParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkContentFilteringParamsWithTimeout(timeout time.Duration) *GetNetworkContentFilteringParams {
	var ()
	return &GetNetworkContentFilteringParams{

		timeout: timeout,
	}
}

// NewGetNetworkContentFilteringParamsWithContext creates a new GetNetworkContentFilteringParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkContentFilteringParamsWithContext(ctx context.Context) *GetNetworkContentFilteringParams {
	var ()
	return &GetNetworkContentFilteringParams{

		Context: ctx,
	}
}

// NewGetNetworkContentFilteringParamsWithHTTPClient creates a new GetNetworkContentFilteringParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkContentFilteringParamsWithHTTPClient(client *http.Client) *GetNetworkContentFilteringParams {
	var ()
	return &GetNetworkContentFilteringParams{
		HTTPClient: client,
	}
}

/*GetNetworkContentFilteringParams contains all the parameters to send to the API endpoint
for the get network content filtering operation typically these are written to a http.Request
*/
type GetNetworkContentFilteringParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network content filtering params
func (o *GetNetworkContentFilteringParams) WithTimeout(timeout time.Duration) *GetNetworkContentFilteringParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network content filtering params
func (o *GetNetworkContentFilteringParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network content filtering params
func (o *GetNetworkContentFilteringParams) WithContext(ctx context.Context) *GetNetworkContentFilteringParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network content filtering params
func (o *GetNetworkContentFilteringParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network content filtering params
func (o *GetNetworkContentFilteringParams) WithHTTPClient(client *http.Client) *GetNetworkContentFilteringParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network content filtering params
func (o *GetNetworkContentFilteringParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network content filtering params
func (o *GetNetworkContentFilteringParams) WithNetworkID(networkID string) *GetNetworkContentFilteringParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network content filtering params
func (o *GetNetworkContentFilteringParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkContentFilteringParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
