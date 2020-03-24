// Code generated by go-swagger; DO NOT EDIT.

package m_x_1_1_n_a_t_rules

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

// NewGetNetworkOneToOneNatRulesParams creates a new GetNetworkOneToOneNatRulesParams object
// with the default values initialized.
func NewGetNetworkOneToOneNatRulesParams() *GetNetworkOneToOneNatRulesParams {
	var ()
	return &GetNetworkOneToOneNatRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkOneToOneNatRulesParamsWithTimeout creates a new GetNetworkOneToOneNatRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkOneToOneNatRulesParamsWithTimeout(timeout time.Duration) *GetNetworkOneToOneNatRulesParams {
	var ()
	return &GetNetworkOneToOneNatRulesParams{

		timeout: timeout,
	}
}

// NewGetNetworkOneToOneNatRulesParamsWithContext creates a new GetNetworkOneToOneNatRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkOneToOneNatRulesParamsWithContext(ctx context.Context) *GetNetworkOneToOneNatRulesParams {
	var ()
	return &GetNetworkOneToOneNatRulesParams{

		Context: ctx,
	}
}

// NewGetNetworkOneToOneNatRulesParamsWithHTTPClient creates a new GetNetworkOneToOneNatRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkOneToOneNatRulesParamsWithHTTPClient(client *http.Client) *GetNetworkOneToOneNatRulesParams {
	var ()
	return &GetNetworkOneToOneNatRulesParams{
		HTTPClient: client,
	}
}

/*GetNetworkOneToOneNatRulesParams contains all the parameters to send to the API endpoint
for the get network one to one nat rules operation typically these are written to a http.Request
*/
type GetNetworkOneToOneNatRulesParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network one to one nat rules params
func (o *GetNetworkOneToOneNatRulesParams) WithTimeout(timeout time.Duration) *GetNetworkOneToOneNatRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network one to one nat rules params
func (o *GetNetworkOneToOneNatRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network one to one nat rules params
func (o *GetNetworkOneToOneNatRulesParams) WithContext(ctx context.Context) *GetNetworkOneToOneNatRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network one to one nat rules params
func (o *GetNetworkOneToOneNatRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network one to one nat rules params
func (o *GetNetworkOneToOneNatRulesParams) WithHTTPClient(client *http.Client) *GetNetworkOneToOneNatRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network one to one nat rules params
func (o *GetNetworkOneToOneNatRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network one to one nat rules params
func (o *GetNetworkOneToOneNatRulesParams) WithNetworkID(networkID string) *GetNetworkOneToOneNatRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network one to one nat rules params
func (o *GetNetworkOneToOneNatRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkOneToOneNatRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
