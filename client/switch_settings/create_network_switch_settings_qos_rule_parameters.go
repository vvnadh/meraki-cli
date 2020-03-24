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

// NewCreateNetworkSwitchSettingsQosRuleParams creates a new CreateNetworkSwitchSettingsQosRuleParams object
// with the default values initialized.
func NewCreateNetworkSwitchSettingsQosRuleParams() *CreateNetworkSwitchSettingsQosRuleParams {
	var ()
	return &CreateNetworkSwitchSettingsQosRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkSwitchSettingsQosRuleParamsWithTimeout creates a new CreateNetworkSwitchSettingsQosRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNetworkSwitchSettingsQosRuleParamsWithTimeout(timeout time.Duration) *CreateNetworkSwitchSettingsQosRuleParams {
	var ()
	return &CreateNetworkSwitchSettingsQosRuleParams{

		timeout: timeout,
	}
}

// NewCreateNetworkSwitchSettingsQosRuleParamsWithContext creates a new CreateNetworkSwitchSettingsQosRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNetworkSwitchSettingsQosRuleParamsWithContext(ctx context.Context) *CreateNetworkSwitchSettingsQosRuleParams {
	var ()
	return &CreateNetworkSwitchSettingsQosRuleParams{

		Context: ctx,
	}
}

// NewCreateNetworkSwitchSettingsQosRuleParamsWithHTTPClient creates a new CreateNetworkSwitchSettingsQosRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNetworkSwitchSettingsQosRuleParamsWithHTTPClient(client *http.Client) *CreateNetworkSwitchSettingsQosRuleParams {
	var ()
	return &CreateNetworkSwitchSettingsQosRuleParams{
		HTTPClient: client,
	}
}

/*CreateNetworkSwitchSettingsQosRuleParams contains all the parameters to send to the API endpoint
for the create network switch settings qos rule operation typically these are written to a http.Request
*/
type CreateNetworkSwitchSettingsQosRuleParams struct {

	/*CreateNetworkSwitchSettingsQosRule*/
	CreateNetworkSwitchSettingsQosRule *models.CreateNetworkSwitchSettingsQosRule
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create network switch settings qos rule params
func (o *CreateNetworkSwitchSettingsQosRuleParams) WithTimeout(timeout time.Duration) *CreateNetworkSwitchSettingsQosRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network switch settings qos rule params
func (o *CreateNetworkSwitchSettingsQosRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network switch settings qos rule params
func (o *CreateNetworkSwitchSettingsQosRuleParams) WithContext(ctx context.Context) *CreateNetworkSwitchSettingsQosRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network switch settings qos rule params
func (o *CreateNetworkSwitchSettingsQosRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network switch settings qos rule params
func (o *CreateNetworkSwitchSettingsQosRuleParams) WithHTTPClient(client *http.Client) *CreateNetworkSwitchSettingsQosRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network switch settings qos rule params
func (o *CreateNetworkSwitchSettingsQosRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkSwitchSettingsQosRule adds the createNetworkSwitchSettingsQosRule to the create network switch settings qos rule params
func (o *CreateNetworkSwitchSettingsQosRuleParams) WithCreateNetworkSwitchSettingsQosRule(createNetworkSwitchSettingsQosRule *models.CreateNetworkSwitchSettingsQosRule) *CreateNetworkSwitchSettingsQosRuleParams {
	o.SetCreateNetworkSwitchSettingsQosRule(createNetworkSwitchSettingsQosRule)
	return o
}

// SetCreateNetworkSwitchSettingsQosRule adds the createNetworkSwitchSettingsQosRule to the create network switch settings qos rule params
func (o *CreateNetworkSwitchSettingsQosRuleParams) SetCreateNetworkSwitchSettingsQosRule(createNetworkSwitchSettingsQosRule *models.CreateNetworkSwitchSettingsQosRule) {
	o.CreateNetworkSwitchSettingsQosRule = createNetworkSwitchSettingsQosRule
}

// WithNetworkID adds the networkID to the create network switch settings qos rule params
func (o *CreateNetworkSwitchSettingsQosRuleParams) WithNetworkID(networkID string) *CreateNetworkSwitchSettingsQosRuleParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network switch settings qos rule params
func (o *CreateNetworkSwitchSettingsQosRuleParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkSwitchSettingsQosRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateNetworkSwitchSettingsQosRule != nil {
		if err := r.SetBodyParam(o.CreateNetworkSwitchSettingsQosRule); err != nil {
			return err
		}
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
