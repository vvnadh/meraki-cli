// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewClaimOrganizationParams creates a new ClaimOrganizationParams object
// with the default values initialized.
func NewClaimOrganizationParams() *ClaimOrganizationParams {
	var ()
	return &ClaimOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewClaimOrganizationParamsWithTimeout creates a new ClaimOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewClaimOrganizationParamsWithTimeout(timeout time.Duration) *ClaimOrganizationParams {
	var ()
	return &ClaimOrganizationParams{

		timeout: timeout,
	}
}

// NewClaimOrganizationParamsWithContext creates a new ClaimOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewClaimOrganizationParamsWithContext(ctx context.Context) *ClaimOrganizationParams {
	var ()
	return &ClaimOrganizationParams{

		Context: ctx,
	}
}

// NewClaimOrganizationParamsWithHTTPClient creates a new ClaimOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewClaimOrganizationParamsWithHTTPClient(client *http.Client) *ClaimOrganizationParams {
	var ()
	return &ClaimOrganizationParams{
		HTTPClient: client,
	}
}

/*ClaimOrganizationParams contains all the parameters to send to the API endpoint
for the claim organization operation typically these are written to a http.Request
*/
type ClaimOrganizationParams struct {

	/*ClaimOrganization*/
	ClaimOrganization *models.ClaimOrganization
	/*OrganizationID*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the claim organization params
func (o *ClaimOrganizationParams) WithTimeout(timeout time.Duration) *ClaimOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the claim organization params
func (o *ClaimOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the claim organization params
func (o *ClaimOrganizationParams) WithContext(ctx context.Context) *ClaimOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the claim organization params
func (o *ClaimOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the claim organization params
func (o *ClaimOrganizationParams) WithHTTPClient(client *http.Client) *ClaimOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the claim organization params
func (o *ClaimOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClaimOrganization adds the claimOrganization to the claim organization params
func (o *ClaimOrganizationParams) WithClaimOrganization(claimOrganization *models.ClaimOrganization) *ClaimOrganizationParams {
	o.SetClaimOrganization(claimOrganization)
	return o
}

// SetClaimOrganization adds the claimOrganization to the claim organization params
func (o *ClaimOrganizationParams) SetClaimOrganization(claimOrganization *models.ClaimOrganization) {
	o.ClaimOrganization = claimOrganization
}

// WithOrganizationID adds the organizationID to the claim organization params
func (o *ClaimOrganizationParams) WithOrganizationID(organizationID string) *ClaimOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the claim organization params
func (o *ClaimOrganizationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *ClaimOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClaimOrganization != nil {
		if err := r.SetBodyParam(o.ClaimOrganization); err != nil {
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
