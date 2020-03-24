// Code generated by go-swagger; DO NOT EDIT.

package security_events

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

// NewGetOrganizationSecurityEventsParams creates a new GetOrganizationSecurityEventsParams object
// with the default values initialized.
func NewGetOrganizationSecurityEventsParams() *GetOrganizationSecurityEventsParams {
	var ()
	return &GetOrganizationSecurityEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationSecurityEventsParamsWithTimeout creates a new GetOrganizationSecurityEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationSecurityEventsParamsWithTimeout(timeout time.Duration) *GetOrganizationSecurityEventsParams {
	var ()
	return &GetOrganizationSecurityEventsParams{

		timeout: timeout,
	}
}

// NewGetOrganizationSecurityEventsParamsWithContext creates a new GetOrganizationSecurityEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationSecurityEventsParamsWithContext(ctx context.Context) *GetOrganizationSecurityEventsParams {
	var ()
	return &GetOrganizationSecurityEventsParams{

		Context: ctx,
	}
}

// NewGetOrganizationSecurityEventsParamsWithHTTPClient creates a new GetOrganizationSecurityEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationSecurityEventsParamsWithHTTPClient(client *http.Client) *GetOrganizationSecurityEventsParams {
	var ()
	return &GetOrganizationSecurityEventsParams{
		HTTPClient: client,
	}
}

/*GetOrganizationSecurityEventsParams contains all the parameters to send to the API endpoint
for the get organization security events operation typically these are written to a http.Request
*/
type GetOrganizationSecurityEventsParams struct {

	/*EndingBefore
	  A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.

	*/
	EndingBefore *string
	/*OrganizationID*/
	OrganizationID string
	/*PerPage
	  The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.

	*/
	PerPage *int32
	/*StartingAfter
	  A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.

	*/
	StartingAfter *string
	/*T0
	  The beginning of the timespan for the data. The maximum lookback period is 365 days from today.

	*/
	T0 *string
	/*T1
	  The end of the timespan for the data. t1 can be a maximum of 365 days after t0.

	*/
	T1 *string
	/*Timespan
	  The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days.

	*/
	Timespan *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) WithTimeout(timeout time.Duration) *GetOrganizationSecurityEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) WithContext(ctx context.Context) *GetOrganizationSecurityEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) WithHTTPClient(client *http.Client) *GetOrganizationSecurityEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) WithEndingBefore(endingBefore *string) *GetOrganizationSecurityEventsParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithOrganizationID adds the organizationID to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) WithOrganizationID(organizationID string) *GetOrganizationSecurityEventsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) WithPerPage(perPage *int32) *GetOrganizationSecurityEventsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) WithStartingAfter(startingAfter *string) *GetOrganizationSecurityEventsParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithT0 adds the t0 to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) WithT0(t0 *string) *GetOrganizationSecurityEventsParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) WithT1(t1 *string) *GetOrganizationSecurityEventsParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) WithTimespan(timespan *float64) *GetOrganizationSecurityEventsParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get organization security events params
func (o *GetOrganizationSecurityEventsParams) SetTimespan(timespan *float64) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationSecurityEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
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

	if o.T0 != nil {

		// query param t0
		var qrT0 string
		if o.T0 != nil {
			qrT0 = *o.T0
		}
		qT0 := qrT0
		if qT0 != "" {
			if err := r.SetQueryParam("t0", qT0); err != nil {
				return err
			}
		}

	}

	if o.T1 != nil {

		// query param t1
		var qrT1 string
		if o.T1 != nil {
			qrT1 = *o.T1
		}
		qT1 := qrT1
		if qT1 != "" {
			if err := r.SetQueryParam("t1", qT1); err != nil {
				return err
			}
		}

	}

	if o.Timespan != nil {

		// query param timespan
		var qrTimespan float64
		if o.Timespan != nil {
			qrTimespan = *o.Timespan
		}
		qTimespan := swag.FormatFloat64(qrTimespan)
		if qTimespan != "" {
			if err := r.SetQueryParam("timespan", qTimespan); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
