// Code generated by go-swagger; DO NOT EDIT.

package switch_settings

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

// NewUpdateNetworkSwitchSettingsDhcpServerPolicyParams creates a new UpdateNetworkSwitchSettingsDhcpServerPolicyParams object
// with the default values initialized.
func NewUpdateNetworkSwitchSettingsDhcpServerPolicyParams() *UpdateNetworkSwitchSettingsDhcpServerPolicyParams {
	var ()
	return &UpdateNetworkSwitchSettingsDhcpServerPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSwitchSettingsDhcpServerPolicyParamsWithTimeout creates a new UpdateNetworkSwitchSettingsDhcpServerPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkSwitchSettingsDhcpServerPolicyParamsWithTimeout(timeout time.Duration) *UpdateNetworkSwitchSettingsDhcpServerPolicyParams {
	var ()
	return &UpdateNetworkSwitchSettingsDhcpServerPolicyParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkSwitchSettingsDhcpServerPolicyParamsWithContext creates a new UpdateNetworkSwitchSettingsDhcpServerPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkSwitchSettingsDhcpServerPolicyParamsWithContext(ctx context.Context) *UpdateNetworkSwitchSettingsDhcpServerPolicyParams {
	var ()
	return &UpdateNetworkSwitchSettingsDhcpServerPolicyParams{

		Context: ctx,
	}
}

// NewUpdateNetworkSwitchSettingsDhcpServerPolicyParamsWithHTTPClient creates a new UpdateNetworkSwitchSettingsDhcpServerPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkSwitchSettingsDhcpServerPolicyParamsWithHTTPClient(client *http.Client) *UpdateNetworkSwitchSettingsDhcpServerPolicyParams {
	var ()
	return &UpdateNetworkSwitchSettingsDhcpServerPolicyParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkSwitchSettingsDhcpServerPolicyParams contains all the parameters to send to the API endpoint
for the update network switch settings dhcp server policy operation typically these are written to a http.Request
*/
type UpdateNetworkSwitchSettingsDhcpServerPolicyParams struct {

	/*NetworkID*/
	NetworkID string
	/*UpdateNetworkSwitchSettingsDhcpServerPolicy*/
	UpdateNetworkSwitchSettingsDhcpServerPolicy *models.UpdateNetworkSwitchSettingsDhcpServerPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network switch settings dhcp server policy params
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) WithTimeout(timeout time.Duration) *UpdateNetworkSwitchSettingsDhcpServerPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network switch settings dhcp server policy params
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network switch settings dhcp server policy params
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) WithContext(ctx context.Context) *UpdateNetworkSwitchSettingsDhcpServerPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network switch settings dhcp server policy params
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network switch settings dhcp server policy params
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) WithHTTPClient(client *http.Client) *UpdateNetworkSwitchSettingsDhcpServerPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network switch settings dhcp server policy params
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network switch settings dhcp server policy params
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) WithNetworkID(networkID string) *UpdateNetworkSwitchSettingsDhcpServerPolicyParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network switch settings dhcp server policy params
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkSwitchSettingsDhcpServerPolicy adds the updateNetworkSwitchSettingsDhcpServerPolicy to the update network switch settings dhcp server policy params
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) WithUpdateNetworkSwitchSettingsDhcpServerPolicy(updateNetworkSwitchSettingsDhcpServerPolicy *models.UpdateNetworkSwitchSettingsDhcpServerPolicy) *UpdateNetworkSwitchSettingsDhcpServerPolicyParams {
	o.SetUpdateNetworkSwitchSettingsDhcpServerPolicy(updateNetworkSwitchSettingsDhcpServerPolicy)
	return o
}

// SetUpdateNetworkSwitchSettingsDhcpServerPolicy adds the updateNetworkSwitchSettingsDhcpServerPolicy to the update network switch settings dhcp server policy params
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) SetUpdateNetworkSwitchSettingsDhcpServerPolicy(updateNetworkSwitchSettingsDhcpServerPolicy *models.UpdateNetworkSwitchSettingsDhcpServerPolicy) {
	o.UpdateNetworkSwitchSettingsDhcpServerPolicy = updateNetworkSwitchSettingsDhcpServerPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSwitchSettingsDhcpServerPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.UpdateNetworkSwitchSettingsDhcpServerPolicy != nil {
		if err := r.SetBodyParam(o.UpdateNetworkSwitchSettingsDhcpServerPolicy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
