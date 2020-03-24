// Code generated by go-swagger; DO NOT EDIT.

package change_log

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

// NewGetOrganizationConfigurationChangesParams creates a new GetOrganizationConfigurationChangesParams object
// with the default values initialized.
func NewGetOrganizationConfigurationChangesParams() *GetOrganizationConfigurationChangesParams {
	var ()
	return &GetOrganizationConfigurationChangesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationConfigurationChangesParamsWithTimeout creates a new GetOrganizationConfigurationChangesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationConfigurationChangesParamsWithTimeout(timeout time.Duration) *GetOrganizationConfigurationChangesParams {
	var ()
	return &GetOrganizationConfigurationChangesParams{

		timeout: timeout,
	}
}

// NewGetOrganizationConfigurationChangesParamsWithContext creates a new GetOrganizationConfigurationChangesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationConfigurationChangesParamsWithContext(ctx context.Context) *GetOrganizationConfigurationChangesParams {
	var ()
	return &GetOrganizationConfigurationChangesParams{

		Context: ctx,
	}
}

// NewGetOrganizationConfigurationChangesParamsWithHTTPClient creates a new GetOrganizationConfigurationChangesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationConfigurationChangesParamsWithHTTPClient(client *http.Client) *GetOrganizationConfigurationChangesParams {
	var ()
	return &GetOrganizationConfigurationChangesParams{
		HTTPClient: client,
	}
}

/*GetOrganizationConfigurationChangesParams contains all the parameters to send to the API endpoint
for the get organization configuration changes operation typically these are written to a http.Request
*/
type GetOrganizationConfigurationChangesParams struct {

	/*AdminID
	  Filters on the given Admin

	*/
	AdminID *string
	/*EndingBefore
	  A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.

	*/
	EndingBefore *string
	/*NetworkID
	  Filters on the given network

	*/
	NetworkID *string
	/*OrganizationID*/
	OrganizationID string
	/*PerPage
	  The number of entries per page returned. Acceptable range is 3 - 5000. Default is 5000.

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
	  The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 365 days.

	*/
	Timespan *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithTimeout(timeout time.Duration) *GetOrganizationConfigurationChangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithContext(ctx context.Context) *GetOrganizationConfigurationChangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithHTTPClient(client *http.Client) *GetOrganizationConfigurationChangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdminID adds the adminID to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithAdminID(adminID *string) *GetOrganizationConfigurationChangesParams {
	o.SetAdminID(adminID)
	return o
}

// SetAdminID adds the adminId to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetAdminID(adminID *string) {
	o.AdminID = adminID
}

// WithEndingBefore adds the endingBefore to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithEndingBefore(endingBefore *string) *GetOrganizationConfigurationChangesParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithNetworkID adds the networkID to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithNetworkID(networkID *string) *GetOrganizationConfigurationChangesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetNetworkID(networkID *string) {
	o.NetworkID = networkID
}

// WithOrganizationID adds the organizationID to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithOrganizationID(organizationID string) *GetOrganizationConfigurationChangesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithPerPage(perPage *int32) *GetOrganizationConfigurationChangesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithStartingAfter(startingAfter *string) *GetOrganizationConfigurationChangesParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithT0 adds the t0 to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithT0(t0 *string) *GetOrganizationConfigurationChangesParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithT1(t1 *string) *GetOrganizationConfigurationChangesParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) WithTimespan(timespan *float64) *GetOrganizationConfigurationChangesParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get organization configuration changes params
func (o *GetOrganizationConfigurationChangesParams) SetTimespan(timespan *float64) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationConfigurationChangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdminID != nil {

		// query param adminId
		var qrAdminID string
		if o.AdminID != nil {
			qrAdminID = *o.AdminID
		}
		qAdminID := qrAdminID
		if qAdminID != "" {
			if err := r.SetQueryParam("adminId", qAdminID); err != nil {
				return err
			}
		}

	}

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

	if o.NetworkID != nil {

		// query param networkId
		var qrNetworkID string
		if o.NetworkID != nil {
			qrNetworkID = *o.NetworkID
		}
		qNetworkID := qrNetworkID
		if qNetworkID != "" {
			if err := r.SetQueryParam("networkId", qNetworkID); err != nil {
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
