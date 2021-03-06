// Code generated by go-swagger; DO NOT EDIT.

package intrusion_settings

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

// NewUpdateNetworkSecurityIntrusionSettingsParams creates a new UpdateNetworkSecurityIntrusionSettingsParams object
// with the default values initialized.
func NewUpdateNetworkSecurityIntrusionSettingsParams() *UpdateNetworkSecurityIntrusionSettingsParams {
	var ()
	return &UpdateNetworkSecurityIntrusionSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSecurityIntrusionSettingsParamsWithTimeout creates a new UpdateNetworkSecurityIntrusionSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkSecurityIntrusionSettingsParamsWithTimeout(timeout time.Duration) *UpdateNetworkSecurityIntrusionSettingsParams {
	var ()
	return &UpdateNetworkSecurityIntrusionSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkSecurityIntrusionSettingsParamsWithContext creates a new UpdateNetworkSecurityIntrusionSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkSecurityIntrusionSettingsParamsWithContext(ctx context.Context) *UpdateNetworkSecurityIntrusionSettingsParams {
	var ()
	return &UpdateNetworkSecurityIntrusionSettingsParams{

		Context: ctx,
	}
}

// NewUpdateNetworkSecurityIntrusionSettingsParamsWithHTTPClient creates a new UpdateNetworkSecurityIntrusionSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkSecurityIntrusionSettingsParamsWithHTTPClient(client *http.Client) *UpdateNetworkSecurityIntrusionSettingsParams {
	var ()
	return &UpdateNetworkSecurityIntrusionSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkSecurityIntrusionSettingsParams contains all the parameters to send to the API endpoint
for the update network security intrusion settings operation typically these are written to a http.Request
*/
type UpdateNetworkSecurityIntrusionSettingsParams struct {

	/*NetworkID*/
	NetworkID string
	/*UpdateNetworkSecurityIntrusionSettings*/
	UpdateNetworkSecurityIntrusionSettings *models.UpdateNetworkSecurityIntrusionSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network security intrusion settings params
func (o *UpdateNetworkSecurityIntrusionSettingsParams) WithTimeout(timeout time.Duration) *UpdateNetworkSecurityIntrusionSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network security intrusion settings params
func (o *UpdateNetworkSecurityIntrusionSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network security intrusion settings params
func (o *UpdateNetworkSecurityIntrusionSettingsParams) WithContext(ctx context.Context) *UpdateNetworkSecurityIntrusionSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network security intrusion settings params
func (o *UpdateNetworkSecurityIntrusionSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network security intrusion settings params
func (o *UpdateNetworkSecurityIntrusionSettingsParams) WithHTTPClient(client *http.Client) *UpdateNetworkSecurityIntrusionSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network security intrusion settings params
func (o *UpdateNetworkSecurityIntrusionSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network security intrusion settings params
func (o *UpdateNetworkSecurityIntrusionSettingsParams) WithNetworkID(networkID string) *UpdateNetworkSecurityIntrusionSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network security intrusion settings params
func (o *UpdateNetworkSecurityIntrusionSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkSecurityIntrusionSettings adds the updateNetworkSecurityIntrusionSettings to the update network security intrusion settings params
func (o *UpdateNetworkSecurityIntrusionSettingsParams) WithUpdateNetworkSecurityIntrusionSettings(updateNetworkSecurityIntrusionSettings *models.UpdateNetworkSecurityIntrusionSettings) *UpdateNetworkSecurityIntrusionSettingsParams {
	o.SetUpdateNetworkSecurityIntrusionSettings(updateNetworkSecurityIntrusionSettings)
	return o
}

// SetUpdateNetworkSecurityIntrusionSettings adds the updateNetworkSecurityIntrusionSettings to the update network security intrusion settings params
func (o *UpdateNetworkSecurityIntrusionSettingsParams) SetUpdateNetworkSecurityIntrusionSettings(updateNetworkSecurityIntrusionSettings *models.UpdateNetworkSecurityIntrusionSettings) {
	o.UpdateNetworkSecurityIntrusionSettings = updateNetworkSecurityIntrusionSettings
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSecurityIntrusionSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.UpdateNetworkSecurityIntrusionSettings != nil {
		if err := r.SetBodyParam(o.UpdateNetworkSecurityIntrusionSettings); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
