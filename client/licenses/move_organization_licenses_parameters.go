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

// NewMoveOrganizationLicensesParams creates a new MoveOrganizationLicensesParams object
// with the default values initialized.
func NewMoveOrganizationLicensesParams() *MoveOrganizationLicensesParams {
	var ()
	return &MoveOrganizationLicensesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMoveOrganizationLicensesParamsWithTimeout creates a new MoveOrganizationLicensesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMoveOrganizationLicensesParamsWithTimeout(timeout time.Duration) *MoveOrganizationLicensesParams {
	var ()
	return &MoveOrganizationLicensesParams{

		timeout: timeout,
	}
}

// NewMoveOrganizationLicensesParamsWithContext creates a new MoveOrganizationLicensesParams object
// with the default values initialized, and the ability to set a context for a request
func NewMoveOrganizationLicensesParamsWithContext(ctx context.Context) *MoveOrganizationLicensesParams {
	var ()
	return &MoveOrganizationLicensesParams{

		Context: ctx,
	}
}

// NewMoveOrganizationLicensesParamsWithHTTPClient creates a new MoveOrganizationLicensesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMoveOrganizationLicensesParamsWithHTTPClient(client *http.Client) *MoveOrganizationLicensesParams {
	var ()
	return &MoveOrganizationLicensesParams{
		HTTPClient: client,
	}
}

/*MoveOrganizationLicensesParams contains all the parameters to send to the API endpoint
for the move organization licenses operation typically these are written to a http.Request
*/
type MoveOrganizationLicensesParams struct {

	/*MoveOrganizationLicenses*/
	MoveOrganizationLicenses *models.MoveOrganizationLicenses
	/*OrganizationID*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the move organization licenses params
func (o *MoveOrganizationLicensesParams) WithTimeout(timeout time.Duration) *MoveOrganizationLicensesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move organization licenses params
func (o *MoveOrganizationLicensesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move organization licenses params
func (o *MoveOrganizationLicensesParams) WithContext(ctx context.Context) *MoveOrganizationLicensesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move organization licenses params
func (o *MoveOrganizationLicensesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move organization licenses params
func (o *MoveOrganizationLicensesParams) WithHTTPClient(client *http.Client) *MoveOrganizationLicensesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move organization licenses params
func (o *MoveOrganizationLicensesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoveOrganizationLicenses adds the moveOrganizationLicenses to the move organization licenses params
func (o *MoveOrganizationLicensesParams) WithMoveOrganizationLicenses(moveOrganizationLicenses *models.MoveOrganizationLicenses) *MoveOrganizationLicensesParams {
	o.SetMoveOrganizationLicenses(moveOrganizationLicenses)
	return o
}

// SetMoveOrganizationLicenses adds the moveOrganizationLicenses to the move organization licenses params
func (o *MoveOrganizationLicensesParams) SetMoveOrganizationLicenses(moveOrganizationLicenses *models.MoveOrganizationLicenses) {
	o.MoveOrganizationLicenses = moveOrganizationLicenses
}

// WithOrganizationID adds the organizationID to the move organization licenses params
func (o *MoveOrganizationLicensesParams) WithOrganizationID(organizationID string) *MoveOrganizationLicensesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the move organization licenses params
func (o *MoveOrganizationLicensesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *MoveOrganizationLicensesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MoveOrganizationLicenses != nil {
		if err := r.SetBodyParam(o.MoveOrganizationLicenses); err != nil {
			return err
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
