// Code generated by go-swagger; DO NOT EDIT.

package licenses

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

// NewRenewOrganizationLicensesSeatsParams creates a new RenewOrganizationLicensesSeatsParams object
// with the default values initialized.
func NewRenewOrganizationLicensesSeatsParams() *RenewOrganizationLicensesSeatsParams {
	var ()
	return &RenewOrganizationLicensesSeatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRenewOrganizationLicensesSeatsParamsWithTimeout creates a new RenewOrganizationLicensesSeatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRenewOrganizationLicensesSeatsParamsWithTimeout(timeout time.Duration) *RenewOrganizationLicensesSeatsParams {
	var ()
	return &RenewOrganizationLicensesSeatsParams{

		timeout: timeout,
	}
}

// NewRenewOrganizationLicensesSeatsParamsWithContext creates a new RenewOrganizationLicensesSeatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRenewOrganizationLicensesSeatsParamsWithContext(ctx context.Context) *RenewOrganizationLicensesSeatsParams {
	var ()
	return &RenewOrganizationLicensesSeatsParams{

		Context: ctx,
	}
}

// NewRenewOrganizationLicensesSeatsParamsWithHTTPClient creates a new RenewOrganizationLicensesSeatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRenewOrganizationLicensesSeatsParamsWithHTTPClient(client *http.Client) *RenewOrganizationLicensesSeatsParams {
	var ()
	return &RenewOrganizationLicensesSeatsParams{
		HTTPClient: client,
	}
}

/*RenewOrganizationLicensesSeatsParams contains all the parameters to send to the API endpoint
for the renew organization licenses seats operation typically these are written to a http.Request
*/
type RenewOrganizationLicensesSeatsParams struct {

	/*OrganizationID*/
	OrganizationID string
	/*RenewOrganizationLicensesSeats*/
	RenewOrganizationLicensesSeats *models.RenewOrganizationLicensesSeats

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the renew organization licenses seats params
func (o *RenewOrganizationLicensesSeatsParams) WithTimeout(timeout time.Duration) *RenewOrganizationLicensesSeatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the renew organization licenses seats params
func (o *RenewOrganizationLicensesSeatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the renew organization licenses seats params
func (o *RenewOrganizationLicensesSeatsParams) WithContext(ctx context.Context) *RenewOrganizationLicensesSeatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the renew organization licenses seats params
func (o *RenewOrganizationLicensesSeatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the renew organization licenses seats params
func (o *RenewOrganizationLicensesSeatsParams) WithHTTPClient(client *http.Client) *RenewOrganizationLicensesSeatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the renew organization licenses seats params
func (o *RenewOrganizationLicensesSeatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the renew organization licenses seats params
func (o *RenewOrganizationLicensesSeatsParams) WithOrganizationID(organizationID string) *RenewOrganizationLicensesSeatsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the renew organization licenses seats params
func (o *RenewOrganizationLicensesSeatsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithRenewOrganizationLicensesSeats adds the renewOrganizationLicensesSeats to the renew organization licenses seats params
func (o *RenewOrganizationLicensesSeatsParams) WithRenewOrganizationLicensesSeats(renewOrganizationLicensesSeats *models.RenewOrganizationLicensesSeats) *RenewOrganizationLicensesSeatsParams {
	o.SetRenewOrganizationLicensesSeats(renewOrganizationLicensesSeats)
	return o
}

// SetRenewOrganizationLicensesSeats adds the renewOrganizationLicensesSeats to the renew organization licenses seats params
func (o *RenewOrganizationLicensesSeatsParams) SetRenewOrganizationLicensesSeats(renewOrganizationLicensesSeats *models.RenewOrganizationLicensesSeats) {
	o.RenewOrganizationLicensesSeats = renewOrganizationLicensesSeats
}

// WriteToRequest writes these params to a swagger request
func (o *RenewOrganizationLicensesSeatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if o.RenewOrganizationLicensesSeats != nil {
		if err := r.SetBodyParam(o.RenewOrganizationLicensesSeats); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
