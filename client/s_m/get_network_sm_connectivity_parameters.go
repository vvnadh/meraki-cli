// Code generated by go-swagger; DO NOT EDIT.

package s_m

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
	"github.com/go-openapi/swag"
)

// NewGetNetworkSmConnectivityParams creates a new GetNetworkSmConnectivityParams object
// with the default values initialized.
func NewGetNetworkSmConnectivityParams() *GetNetworkSmConnectivityParams {
	var ()
	return &GetNetworkSmConnectivityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmConnectivityParamsWithTimeout creates a new GetNetworkSmConnectivityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkSmConnectivityParamsWithTimeout(timeout time.Duration) *GetNetworkSmConnectivityParams {
	var ()
	return &GetNetworkSmConnectivityParams{

		timeout: timeout,
	}
}

// NewGetNetworkSmConnectivityParamsWithContext creates a new GetNetworkSmConnectivityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkSmConnectivityParamsWithContext(ctx context.Context) *GetNetworkSmConnectivityParams {
	var ()
	return &GetNetworkSmConnectivityParams{

		Context: ctx,
	}
}

// NewGetNetworkSmConnectivityParamsWithHTTPClient creates a new GetNetworkSmConnectivityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkSmConnectivityParamsWithHTTPClient(client *http.Client) *GetNetworkSmConnectivityParams {
	var ()
	return &GetNetworkSmConnectivityParams{
		HTTPClient: client,
	}
}

/*GetNetworkSmConnectivityParams contains all the parameters to send to the API endpoint
for the get network sm connectivity operation typically these are written to a http.Request
*/
type GetNetworkSmConnectivityParams struct {

	/*EndingBefore
	  A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.

	*/
	EndingBefore *string
	/*ID*/
	ID string
	/*NetworkID*/
	NetworkID string
	/*PerPage
	  The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.

	*/
	PerPage *int32
	/*StartingAfter
	  A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.

	*/
	StartingAfter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) WithTimeout(timeout time.Duration) *GetNetworkSmConnectivityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) WithContext(ctx context.Context) *GetNetworkSmConnectivityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) WithHTTPClient(client *http.Client) *GetNetworkSmConnectivityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) WithEndingBefore(endingBefore *string) *GetNetworkSmConnectivityParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithID adds the id to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) WithID(id string) *GetNetworkSmConnectivityParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) SetID(id string) {
	o.ID = id
}

// WithNetworkID adds the networkID to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) WithNetworkID(networkID string) *GetNetworkSmConnectivityParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPerPage adds the perPage to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) WithPerPage(perPage *int32) *GetNetworkSmConnectivityParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) WithStartingAfter(startingAfter *string) *GetNetworkSmConnectivityParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get network sm connectivity params
func (o *GetNetworkSmConnectivityParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmConnectivityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndingBefore != nil {

		// query param endingBefore
		var qrEndingBefore string
		if o.EndingBefore != nil {
			qrEndingBefore = *o.EndingBefore
		}
		qEndingBefore := qrEndingBefore
		if qEndingBefore != "" {
			if err := r.SetQueryParam("endingBefore", qEndingBefore); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.PerPage != nil {

		// query param perPage
		var qrPerPage int32
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("perPage", qPerPage); err != nil {
				return err
			}
		}

	}

	if o.StartingAfter != nil {

		// query param startingAfter
		var qrStartingAfter string
		if o.StartingAfter != nil {
			qrStartingAfter = *o.StartingAfter
		}
		qStartingAfter := qrStartingAfter
		if qStartingAfter != "" {
			if err := r.SetQueryParam("startingAfter", qStartingAfter); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
