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

// NewSplitNetworkParams creates a new SplitNetworkParams object
// with the default values initialized.
func NewSplitNetworkParams() *SplitNetworkParams {
	var ()
	return &SplitNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSplitNetworkParamsWithTimeout creates a new SplitNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSplitNetworkParamsWithTimeout(timeout time.Duration) *SplitNetworkParams {
	var ()
	return &SplitNetworkParams{

		timeout: timeout,
	}
}

// NewSplitNetworkParamsWithContext creates a new SplitNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewSplitNetworkParamsWithContext(ctx context.Context) *SplitNetworkParams {
	var ()
	return &SplitNetworkParams{

		Context: ctx,
	}
}

// NewSplitNetworkParamsWithHTTPClient creates a new SplitNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSplitNetworkParamsWithHTTPClient(client *http.Client) *SplitNetworkParams {
	var ()
	return &SplitNetworkParams{
		HTTPClient: client,
	}
}

/*SplitNetworkParams contains all the parameters to send to the API endpoint
for the split network operation typically these are written to a http.Request
*/
type SplitNetworkParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the split network params
func (o *SplitNetworkParams) WithTimeout(timeout time.Duration) *SplitNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the split network params
func (o *SplitNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the split network params
func (o *SplitNetworkParams) WithContext(ctx context.Context) *SplitNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the split network params
func (o *SplitNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the split network params
func (o *SplitNetworkParams) WithHTTPClient(client *http.Client) *SplitNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the split network params
func (o *SplitNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the split network params
func (o *SplitNetworkParams) WithNetworkID(networkID string) *SplitNetworkParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the split network params
func (o *SplitNetworkParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *SplitNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
