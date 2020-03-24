// Code generated by go-swagger; DO NOT EDIT.

package s_a_m_l_roles

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

// NewGetOrganizationSamlRoleParams creates a new GetOrganizationSamlRoleParams object
// with the default values initialized.
func NewGetOrganizationSamlRoleParams() *GetOrganizationSamlRoleParams {
	var ()
	return &GetOrganizationSamlRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationSamlRoleParamsWithTimeout creates a new GetOrganizationSamlRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationSamlRoleParamsWithTimeout(timeout time.Duration) *GetOrganizationSamlRoleParams {
	var ()
	return &GetOrganizationSamlRoleParams{

		timeout: timeout,
	}
}

// NewGetOrganizationSamlRoleParamsWithContext creates a new GetOrganizationSamlRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationSamlRoleParamsWithContext(ctx context.Context) *GetOrganizationSamlRoleParams {
	var ()
	return &GetOrganizationSamlRoleParams{

		Context: ctx,
	}
}

// NewGetOrganizationSamlRoleParamsWithHTTPClient creates a new GetOrganizationSamlRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationSamlRoleParamsWithHTTPClient(client *http.Client) *GetOrganizationSamlRoleParams {
	var ()
	return &GetOrganizationSamlRoleParams{
		HTTPClient: client,
	}
}

/*GetOrganizationSamlRoleParams contains all the parameters to send to the API endpoint
for the get organization saml role operation typically these are written to a http.Request
*/
type GetOrganizationSamlRoleParams struct {

	/*OrganizationID*/
	OrganizationID string
	/*SamlRoleID*/
	SamlRoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organization saml role params
func (o *GetOrganizationSamlRoleParams) WithTimeout(timeout time.Duration) *GetOrganizationSamlRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization saml role params
func (o *GetOrganizationSamlRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization saml role params
func (o *GetOrganizationSamlRoleParams) WithContext(ctx context.Context) *GetOrganizationSamlRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization saml role params
func (o *GetOrganizationSamlRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization saml role params
func (o *GetOrganizationSamlRoleParams) WithHTTPClient(client *http.Client) *GetOrganizationSamlRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization saml role params
func (o *GetOrganizationSamlRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization saml role params
func (o *GetOrganizationSamlRoleParams) WithOrganizationID(organizationID string) *GetOrganizationSamlRoleParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization saml role params
func (o *GetOrganizationSamlRoleParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithSamlRoleID adds the samlRoleID to the get organization saml role params
func (o *GetOrganizationSamlRoleParams) WithSamlRoleID(samlRoleID string) *GetOrganizationSamlRoleParams {
	o.SetSamlRoleID(samlRoleID)
	return o
}

// SetSamlRoleID adds the samlRoleId to the get organization saml role params
func (o *GetOrganizationSamlRoleParams) SetSamlRoleID(samlRoleID string) {
	o.SamlRoleID = samlRoleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationSamlRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	// path param samlRoleId
	if err := r.SetPathParam("samlRoleId", o.SamlRoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
