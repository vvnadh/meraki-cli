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

// NewUpdateOrganizationThirdPartyVPNPeersParams creates a new UpdateOrganizationThirdPartyVPNPeersParams object
// with the default values initialized.
func NewUpdateOrganizationThirdPartyVPNPeersParams() *UpdateOrganizationThirdPartyVPNPeersParams {
	var ()
	return &UpdateOrganizationThirdPartyVPNPeersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationThirdPartyVPNPeersParamsWithTimeout creates a new UpdateOrganizationThirdPartyVPNPeersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOrganizationThirdPartyVPNPeersParamsWithTimeout(timeout time.Duration) *UpdateOrganizationThirdPartyVPNPeersParams {
	var ()
	return &UpdateOrganizationThirdPartyVPNPeersParams{

		timeout: timeout,
	}
}

// NewUpdateOrganizationThirdPartyVPNPeersParamsWithContext creates a new UpdateOrganizationThirdPartyVPNPeersParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOrganizationThirdPartyVPNPeersParamsWithContext(ctx context.Context) *UpdateOrganizationThirdPartyVPNPeersParams {
	var ()
	return &UpdateOrganizationThirdPartyVPNPeersParams{

		Context: ctx,
	}
}

// NewUpdateOrganizationThirdPartyVPNPeersParamsWithHTTPClient creates a new UpdateOrganizationThirdPartyVPNPeersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOrganizationThirdPartyVPNPeersParamsWithHTTPClient(client *http.Client) *UpdateOrganizationThirdPartyVPNPeersParams {
	var ()
	return &UpdateOrganizationThirdPartyVPNPeersParams{
		HTTPClient: client,
	}
}

/*UpdateOrganizationThirdPartyVPNPeersParams contains all the parameters to send to the API endpoint
for the update organization third party v p n peers operation typically these are written to a http.Request
*/
type UpdateOrganizationThirdPartyVPNPeersParams struct {

	/*OrganizationID*/
	OrganizationID string
	/*UpdateOrganizationThirdPartyVPNPeers*/
	UpdateOrganizationThirdPartyVPNPeers *models.UpdateOrganizationThirdPartyVPNPeers

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update organization third party v p n peers params
func (o *UpdateOrganizationThirdPartyVPNPeersParams) WithTimeout(timeout time.Duration) *UpdateOrganizationThirdPartyVPNPeersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization third party v p n peers params
func (o *UpdateOrganizationThirdPartyVPNPeersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization third party v p n peers params
func (o *UpdateOrganizationThirdPartyVPNPeersParams) WithContext(ctx context.Context) *UpdateOrganizationThirdPartyVPNPeersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization third party v p n peers params
func (o *UpdateOrganizationThirdPartyVPNPeersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization third party v p n peers params
func (o *UpdateOrganizationThirdPartyVPNPeersParams) WithHTTPClient(client *http.Client) *UpdateOrganizationThirdPartyVPNPeersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization third party v p n peers params
func (o *UpdateOrganizationThirdPartyVPNPeersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the update organization third party v p n peers params
func (o *UpdateOrganizationThirdPartyVPNPeersParams) WithOrganizationID(organizationID string) *UpdateOrganizationThirdPartyVPNPeersParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization third party v p n peers params
func (o *UpdateOrganizationThirdPartyVPNPeersParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUpdateOrganizationThirdPartyVPNPeers adds the updateOrganizationThirdPartyVPNPeers to the update organization third party v p n peers params
func (o *UpdateOrganizationThirdPartyVPNPeersParams) WithUpdateOrganizationThirdPartyVPNPeers(updateOrganizationThirdPartyVPNPeers *models.UpdateOrganizationThirdPartyVPNPeers) *UpdateOrganizationThirdPartyVPNPeersParams {
	o.SetUpdateOrganizationThirdPartyVPNPeers(updateOrganizationThirdPartyVPNPeers)
	return o
}

// SetUpdateOrganizationThirdPartyVPNPeers adds the updateOrganizationThirdPartyVPNPeers to the update organization third party v p n peers params
func (o *UpdateOrganizationThirdPartyVPNPeersParams) SetUpdateOrganizationThirdPartyVPNPeers(updateOrganizationThirdPartyVPNPeers *models.UpdateOrganizationThirdPartyVPNPeers) {
	o.UpdateOrganizationThirdPartyVPNPeers = updateOrganizationThirdPartyVPNPeers
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationThirdPartyVPNPeersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if o.UpdateOrganizationThirdPartyVPNPeers != nil {
		if err := r.SetBodyParam(o.UpdateOrganizationThirdPartyVPNPeers); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
