// Code generated by go-swagger; DO NOT EDIT.

package traffic_shaping

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

// NewGetNetworkTrafficShapingDscpTaggingOptionsParams creates a new GetNetworkTrafficShapingDscpTaggingOptionsParams object
// with the default values initialized.
func NewGetNetworkTrafficShapingDscpTaggingOptionsParams() *GetNetworkTrafficShapingDscpTaggingOptionsParams {
	var ()
	return &GetNetworkTrafficShapingDscpTaggingOptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkTrafficShapingDscpTaggingOptionsParamsWithTimeout creates a new GetNetworkTrafficShapingDscpTaggingOptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkTrafficShapingDscpTaggingOptionsParamsWithTimeout(timeout time.Duration) *GetNetworkTrafficShapingDscpTaggingOptionsParams {
	var ()
	return &GetNetworkTrafficShapingDscpTaggingOptionsParams{

		timeout: timeout,
	}
}

// NewGetNetworkTrafficShapingDscpTaggingOptionsParamsWithContext creates a new GetNetworkTrafficShapingDscpTaggingOptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkTrafficShapingDscpTaggingOptionsParamsWithContext(ctx context.Context) *GetNetworkTrafficShapingDscpTaggingOptionsParams {
	var ()
	return &GetNetworkTrafficShapingDscpTaggingOptionsParams{

		Context: ctx,
	}
}

// NewGetNetworkTrafficShapingDscpTaggingOptionsParamsWithHTTPClient creates a new GetNetworkTrafficShapingDscpTaggingOptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkTrafficShapingDscpTaggingOptionsParamsWithHTTPClient(client *http.Client) *GetNetworkTrafficShapingDscpTaggingOptionsParams {
	var ()
	return &GetNetworkTrafficShapingDscpTaggingOptionsParams{
		HTTPClient: client,
	}
}

/*GetNetworkTrafficShapingDscpTaggingOptionsParams contains all the parameters to send to the API endpoint
for the get network traffic shaping dscp tagging options operation typically these are written to a http.Request
*/
type GetNetworkTrafficShapingDscpTaggingOptionsParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network traffic shaping dscp tagging options params
func (o *GetNetworkTrafficShapingDscpTaggingOptionsParams) WithTimeout(timeout time.Duration) *GetNetworkTrafficShapingDscpTaggingOptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network traffic shaping dscp tagging options params
func (o *GetNetworkTrafficShapingDscpTaggingOptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network traffic shaping dscp tagging options params
func (o *GetNetworkTrafficShapingDscpTaggingOptionsParams) WithContext(ctx context.Context) *GetNetworkTrafficShapingDscpTaggingOptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network traffic shaping dscp tagging options params
func (o *GetNetworkTrafficShapingDscpTaggingOptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network traffic shaping dscp tagging options params
func (o *GetNetworkTrafficShapingDscpTaggingOptionsParams) WithHTTPClient(client *http.Client) *GetNetworkTrafficShapingDscpTaggingOptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network traffic shaping dscp tagging options params
func (o *GetNetworkTrafficShapingDscpTaggingOptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network traffic shaping dscp tagging options params
func (o *GetNetworkTrafficShapingDscpTaggingOptionsParams) WithNetworkID(networkID string) *GetNetworkTrafficShapingDscpTaggingOptionsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network traffic shaping dscp tagging options params
func (o *GetNetworkTrafficShapingDscpTaggingOptionsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkTrafficShapingDscpTaggingOptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
